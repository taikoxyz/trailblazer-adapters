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

var (
	logTransferSigHash = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	logStakeSigHash    = crypto.Keccak256Hash([]byte("Stake(address,uint256,address)"))
)

// TransferIndexer is an implementation of LogsIndexer for ERC20 transfer logs.
type TransferIndexer struct {
	tokens    []common.Address
	whitelist map[string]struct{}
}

// NewTransferIndexer creates a new TransferIndexer.
func NewTransferIndexer(tokens []common.Address, whitelist map[string]struct{}) *TransferIndexer {
	return &TransferIndexer{
		tokens:    tokens,
		whitelist: whitelist,
	}
}

func (indexer *TransferIndexer) Address() []common.Address {
	return indexer.tokens
}

// IndexLogs processes logs for ERC20 transfers.
func (indexer *TransferIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.LPTransfer, error) {
	var result []adapters.LPTransfer
	for _, vLog := range logs {
		if !isERC20Transfer(vLog) {
			continue
		}
		transferData, err := indexer.ProcessLog(ctx, chainID, client, vLog)
		if err != nil {
			return nil, err
		}
		if transferData != nil {
			result = append(result, *transferData)
		}
	}
	return result, nil
}

func isERC20Transfer(vLog types.Log) bool {
	return len(vLog.Topics) == 3 && vLog.Topics[0].Hex() == logTransferSigHash.Hex()
}

func (indexer *TransferIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.LPTransfer, error) {
	// Extract "from" and "to" addresses from the log
	to := common.BytesToAddress(vLog.Topics[2].Bytes()[12:])
	_, exists := indexer.whitelist[to.Hex()]
	if !exists {
		return nil, nil
	}
	txReceipt, err := client.TransactionReceipt(ctx, vLog.TxHash)
	if err != nil {
		return nil, err
	}
	from := adapters.ZeroAddress
	for _, log := range txReceipt.Logs {
		if log.Topics[0].Hex() == logStakeSigHash.Hex() {
			from = common.BytesToAddress(log.Topics[2].Bytes()[12:])
		}
	}

	// Unpack the transfer event
	var transferEvent struct {
		Value *big.Int
	}
	if err := unpackTransferEvent(vLog, &transferEvent); err != nil {
		return nil, err
	}

	// Fetch the block details
	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	// Initialize the LP token caller
	token, err := ritsu.NewRitsuCaller(vLog.Address, client)
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
	_, token0Decimals, err := fetchTokenDetails(token0Address, client)
	if err != nil {
		return nil, err
	}

	_, token1Decimals, err := fetchTokenDetails(token1Address, client)
	if err != nil {
		return nil, err
	}

	// Calculate the user's share of each token in the pool
	token0Share := calculateTokenShare(shareOfPool, reserves.Reserve0)
	token1Share := calculateTokenShare(shareOfPool, reserves.Reserve1)

	// Return the LPTransfer struct with calculated values
	return &adapters.LPTransfer{
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
	}, nil
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
