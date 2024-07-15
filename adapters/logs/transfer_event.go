package logs

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/erc20"
)

var _ adapters.TransferLogsIndexer = (*TransferIndexer)(nil)

// TransferIndexer is an implementation of LogsIndexer for ERC20 transfer logs.
type TransferIndexer struct {
	Addresses []common.Address
}

// NewTransferIndexer creates a new TransferIndexer.
func NewTransferIndexer(addresses []common.Address) *TransferIndexer {
	return &TransferIndexer{Addresses: addresses}
}

// IndexLogs processes logs for ERC20 transfers.
func (indexer *TransferIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) (*[]adapters.TransferData, error) {
	var result []adapters.TransferData
	for _, vLog := range logs {
		if !indexer.isERC20Transfer(vLog) {
			continue
		}
		transferData, err := indexer.ProcessLog(ctx, chainID, client, vLog)
		if err != nil {
			return nil, err
		}
		result = append(result, *transferData)
	}
	return &result, nil
}

func (indexer *TransferIndexer) isERC20Transfer(vLog types.Log) bool {
	if len(vLog.Topics) == 0 {
		return false
	}
	if vLog.Topics[0].Hex() != "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" {
		return false
	}

	if len(vLog.Topics) == 3 {
		return true
	}

	return false
}

// ProcessLog processes a single ERC20 transfer log.
func (indexer *TransferIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.TransferData, error) {
	to := fmt.Sprintf("0x%v", common.Bytes2Hex(vLog.Topics[2].Bytes()[12:]))
	from := fmt.Sprintf("0x%v", common.Bytes2Hex(vLog.Topics[1].Bytes()[12:]))
	toHex := common.HexToAddress(to)
	fromHex := common.HexToAddress(from)

	type TransferEvent struct {
		Value *big.Int
	}

	var t TransferEvent
	erc20ABI, err := abi.JSON(strings.NewReader(erc20.ABI))
	if err != nil {
		return nil, err
	}
	err = erc20ABI.UnpackIntoInterface(&t, "Transfer", vLog.Data)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	return &adapters.TransferData{
		From:  fromHex,
		To:    toHex,
		Time:  block.Time(),
		Value: t.Value,
	}, nil
}
