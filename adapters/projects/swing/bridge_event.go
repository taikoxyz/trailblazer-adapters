package swing

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/swing"
)

var (
	logCrossChainSwapSigHash = crypto.Keccak256Hash([]byte("CrosschainSwapRequest(bytes32,bytes32,bytes32,address,address,address,address,uint256,uint256,uint256,uint8)"))
)

// TransferIndexer is an implementation of LogsIndexer for ERC20 transfer logs.
type CrosschainSwapRequestIndexer struct {
	TargetAddresses []common.Address
}

func NewCrosschainSwapRequestIndexer() *CrosschainSwapRequestIndexer {
	return &CrosschainSwapRequestIndexer{TargetAddresses: []common.Address{
		common.HexToAddress("0x42df81c742CAe6F6D91E136b1AA5C7e14CB394FB"), // SwitchEvent by Swing
	}}
}

func (indexer *CrosschainSwapRequestIndexer) Addresses() []common.Address {
	return indexer.TargetAddresses
}

type CrosschainSwapRequestEvent struct {
	Id               [32]byte
	BridgeTransferId [32]byte
	Bridge           [32]byte
	From             common.Address
	FromToken        common.Address
	BridgeToken      common.Address
	DestToken        common.Address
	FromAmount       *big.Int
	BridgeAmount     *big.Int
	DstAmount        *big.Int
	Status           uint8
}

// IndexLogs processes logs for ERC20 transfers.
func (indexer *CrosschainSwapRequestIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.Whitelist, error) {
	var result []adapters.Whitelist
	for _, vLog := range logs {
		if !indexer.isRelevantLog(vLog.Topics[0]) {
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

func (indexer *CrosschainSwapRequestIndexer) isRelevantLog(topic common.Hash) bool {
	return topic.Hex() == logCrossChainSwapSigHash.Hex()
}

func (indexer *CrosschainSwapRequestIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.Whitelist, error) {
	var newCrossChainSwapEvent CrosschainSwapRequestEvent

	crossChainSwapRequestABI, err := abi.JSON(strings.NewReader(swing.ABI))
	if err != nil {
		return nil, err
	}

	err = crossChainSwapRequestABI.UnpackIntoInterface(&newCrossChainSwapEvent, "CrosschainSwapRequest", vLog.Data)
	if err != nil {
		return nil, err
	}

	txn, err := client.TransactionInBlock(ctx, vLog.BlockHash, vLog.TxIndex)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	sender, err := client.TransactionSender(ctx, txn, vLog.BlockHash, vLog.TxIndex)
	if err != nil {
		return nil, err
	}

	return &adapters.Whitelist{
		User:        sender,
		Time:        block.Time(),
		BlockNumber: block.Number().Uint64(),
	}, nil
}
