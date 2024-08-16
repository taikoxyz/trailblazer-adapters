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
)

// TransferIndexer is an implementation of LogsIndexer for ERC20 transfer logs.
type TransferIndexer struct {
	token common.Address
}

// NewTransferIndexer creates a new TransferIndexer.
func NewTransferIndexer(token common.Address) *TransferIndexer {
	return &TransferIndexer{
		token: token,
	}
}

func (indexer *TransferIndexer) Address() common.Address {
	return indexer.token
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
		result = append(result, *transferData)
	}
	return result, nil
}

func isERC20Transfer(vLog types.Log) bool {
	return len(vLog.Topics) == 3 && vLog.Topics[0].Hex() == logTransferSigHash.Hex()
}

func (indexer *TransferIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.LPTransfer, error) {
	to := common.BytesToAddress(vLog.Topics[2].Bytes()[12:])
	from := common.BytesToAddress(vLog.Topics[1].Bytes()[12:])

	var transferEvent struct {
		Value *big.Int
	}

	ritsuABI, err := abi.JSON(strings.NewReader(erc20.Erc20ABI))
	if err != nil {
		return nil, err
	}

	err = ritsuABI.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	token, err := ritsu.NewRitsuCaller(indexer.token, client)
	if err != nil {
		return nil, err
	}

	reserves, err := token.GetReserves(nil)
	if err != nil {
		return nil, err
	}

	token0Address, err := token.Token0(&bind.CallOpts{
		BlockNumber: block.Number(),
	})
	if err != nil {
		return nil, err
	}

	token1Address, err := token.Token1(&bind.CallOpts{
		BlockNumber: block.Number(),
	})
	if err != nil {
		return nil, err
	}

	token0, err := erc20.NewErc20Caller(token0Address, client)
	if err != nil {
		return nil, err
	}

	token1, err := erc20.NewErc20Caller(token1Address, client)
	if err != nil {
		return nil, err
	}

	totalSupply, err := token.TotalSupply(&bind.CallOpts{
		BlockNumber: block.Number(),
	})
	if err != nil {
		return nil, err
	}
	// Calculate user's share of the pool as a fraction
	shareOfPool := new(big.Rat).SetFrac(transferEvent.Value, totalSupply)

	token0Decimals, err := token0.Decimals(nil)
	if err != nil {
		return nil, err
	}

	token1Decimals, err := token1.Decimals(nil)
	if err != nil {
		return nil, err
	}

	// Calculate the user's share of each token in the pool
	token0Share := new(big.Int).Mul(shareOfPool.Num(), reserves.Reserve0)
	token0Share.Div(token0Share, shareOfPool.Denom())

	token1Share := new(big.Int).Mul(shareOfPool.Num(), reserves.Reserve1)
	token1Share.Div(token1Share, shareOfPool.Denom())

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
