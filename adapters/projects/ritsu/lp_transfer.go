package ritsu

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/erc20"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/ritsu"
)

const (
	logTransferSignature string = "Transfer(address,address,uint256)"
	logStakeSignature    string = "Stake(address,uint256,address)"
)

func LPAdresses() []string {
	return []string{
		"0x7c38E9389B27668280E5aaAc372eBCb2ECc1c5E0", // https://taikoscan.io/address/0x7c38E9389B27668280E5aaAc372eBCb2ECc1c5E0
		"0x6c7839E0CE8AdA360a865E18a111A462d08DC15a", // https://taikoscan.io/address/0x6c7839E0CE8AdA360a865E18a111A462d08DC15a
	}
}

func Whitelist() []string {
	return []string{
		"0xaE7850cBbA0d7303eD06661c9B06f4A5127Ca75D", // https://taikoscan.io/address/0xaE7850cBbA0d7303eD06661c9B06f4A5127Ca75D
		"0xd7506E9Ebac0Ff14Ce3C600DC6Fd94240e5612D4", // https://taikoscan.io/address/0xd7506E9Ebac0Ff14Ce3C600DC6Fd94240e5612D4
	}
}

type TransferIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
	Whitelist map[string]struct{}
}

func NewTransferIndexer(client *ethclient.Client, addresses []common.Address, whitelist []string) *TransferIndexer {
	indexer := &TransferIndexer{
		client:    client,
		addresses: addresses,
		Whitelist: map[string]struct{}{},
	}

	for _, addr := range whitelist {
		indexer.Whitelist[addr] = struct{}{}
	}

	return indexer
}

var _ adapters.LogIndexer[adapters.LPTransfer] = &TransferIndexer{}

func (indexer *TransferIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *TransferIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.LPTransfer, error) {
	var lpTransfers []adapters.LPTransfer

	for _, l := range logs {
		if !indexer.isERC20Transfer(l) {
			continue
		}

		// Extract "from" and "to" addresses from the log
		to := common.BytesToAddress(l.Topics[2].Bytes()[12:])
		_, exists := indexer.Whitelist[to.Hex()]
		if !exists {
			return nil, nil
		}
		txReceipt, err := indexer.client.TransactionReceipt(ctx, l.TxHash)
		if err != nil {
			return nil, err
		}
		from := adapters.ZeroAddress()
		for _, log := range txReceipt.Logs {
			if log.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logStakeSignature)).Hex() {
				from = common.BytesToAddress(log.Topics[2].Bytes()[12:])
			}
		}

		// Unpack the transfer event
		var transferEvent struct {
			Value *big.Int
		}
		if err := unpackTransferEvent(l, &transferEvent); err != nil {
			return nil, err
		}

		// Fetch the block details
		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		// Initialize the LP token caller
		token, err := ritsu.NewRitsuCaller(l.Address, indexer.client)
		if err != nil {
			return nil, err
		}

		// Fetch reserve balances and token addresses
		reserves, token0Address, token1Address, err := fetchReservesAndTokens(token, block.Number())
		if err != nil {
			return nil, err
		}

		// Calculate user's share of the pool
		shareOfPool, err := calculateShareOfPool(transferEvent.Value, token, block.Number())
		if err != nil {
			return nil, err
		}

		// Fetch token details (decimals)
		_, token0Decimals, err := fetchTokenDetails(token0Address, indexer.client)
		if err != nil {
			return nil, err
		}

		_, token1Decimals, err := fetchTokenDetails(token1Address, indexer.client)
		if err != nil {
			return nil, err
		}

		// Calculate the user's share of each token in the pool
		token0Share := calculateTokenShare(shareOfPool, reserves.Reserve0)
		token1Share := calculateTokenShare(shareOfPool, reserves.Reserve1)

		// Return the LPTransfer struct with calculated values
		lpTransfer := &adapters.LPTransfer{
			From:           from,
			To:             to,
			Token0Amount:   token0Share,
			Token0Decimals: token0Decimals,
			Token0:         token0Address,
			Token1Amount:   token1Share,
			Token1Decimals: token1Decimals,
			Token1:         token1Address,
			Time:           block.Time(),
			BlockNumber:    block.Number().Uint64(),
			TxHash:         l.TxHash,
		}

		lpTransfers = append(lpTransfers, *lpTransfer)
	}

	return lpTransfers, nil
}

func (indexer *TransferIndexer) isERC20Transfer(l types.Log) bool {
	return len(l.Topics) == 3 && l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logTransferSignature)).Hex()
}

// Helper function to unpack the transfer event from the log
func unpackTransferEvent(vLog types.Log, transferEvent *struct{ Value *big.Int }) error {
	ritsuABI, err := abi.JSON(strings.NewReader(erc20.Erc20ABI))
	if err != nil {
		return err
	}
	return ritsuABI.UnpackIntoInterface(transferEvent, "Transfer", vLog.Data)
}

// Helper function to fetch reserves and token addresses from the LP contract
func fetchReservesAndTokens(token *ritsu.RitsuCaller, blockNumber *big.Int) (reserves struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, token0Address, token1Address common.Address, err error) {
	reserves, err = token.GetReserves(nil)
	if err != nil {
		return
	}

	token0Address, err = token.Token0(&bind.CallOpts{BlockNumber: blockNumber})
	if err != nil {
		return
	}

	token1Address, err = token.Token1(&bind.CallOpts{BlockNumber: blockNumber})
	return
}

// Helper function to calculate the user's share of the pool
func calculateShareOfPool(transferValue *big.Int, token *ritsu.RitsuCaller, blockNumber *big.Int) (*big.Rat, error) {
	totalSupply, err := token.TotalSupply(&bind.CallOpts{BlockNumber: blockNumber})
	if err != nil {
		return nil, err
	}
	return new(big.Rat).SetFrac(transferValue, totalSupply), nil
}

// Helper function to fetch token details (caller and decimals)
func fetchTokenDetails(tokenAddress common.Address, client *ethclient.Client) (token *erc20.Erc20Caller, decimals uint8, err error) {
	token, err = erc20.NewErc20Caller(tokenAddress, client)
	if err != nil {
		return
	}

	decimals, err = token.Decimals(nil)
	return
}

// Helper function to calculate the share of a specific token in the pool
func calculateTokenShare(shareOfPool *big.Rat, reserve *big.Int) *big.Int {
	tokenShare := new(big.Int).Mul(shareOfPool.Num(), reserve)
	return tokenShare.Div(tokenShare, shareOfPool.Denom())
}
