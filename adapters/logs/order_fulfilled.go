package logs

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
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/order"
)

var _ adapters.TransferLogsIndexer = (*OrderFulfilledIndexer)(nil)

var (
	logOrderFulfilledSigHash = crypto.Keccak256Hash([]byte("OrderFulfilled(bytes32,address,address,address,(uint8,address,uint256,uint256)[],(uint8,address,uint256,uint256,address)[])"))
)

type OrderFulfilledIndexer struct {
	Addresses []common.Address
}

type SpentItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
}

type ReceivedItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
	Recipient  common.Address
}

type OrderFulfilledEvent struct {
	OrderHash     [32]byte
	Offerer       common.Address
	Zone          common.Address
	Recipient     common.Address
	Offer         []SpentItem
	Consideration []ReceivedItem
}

func NewOrderFulfilledIndexer() *OrderFulfilledIndexer {
	return &OrderFulfilledIndexer{Addresses: []common.Address{
		common.HexToAddress("0x00000000000000ADc04C56Bf30aC9d3c0aAF14dC"),
	}}
}

func (indexer *OrderFulfilledIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.TransferData, error) {
	var result []adapters.TransferData
	for _, vLog := range logs {
		if !indexer.isRelevantLog(vLog.Topics[0]) {
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

func (indexer *OrderFulfilledIndexer) isRelevantLog(topic common.Hash) bool {
	return topic.Hex() == logOrderFulfilledSigHash.Hex()
}

func (indexer *OrderFulfilledIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.TransferData, error) {
	var orderFulfilledEvent OrderFulfilledEvent

	orderFulfilledABI, err := abi.JSON(strings.NewReader(order.ABI))
	if err != nil {
		return nil, err
	}
	err = orderFulfilledABI.UnpackIntoInterface(&orderFulfilledEvent, "OrderFulfilled", vLog.Data)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	return &adapters.TransferData{
		To:          orderFulfilledEvent.Recipient,
		Time:        block.Time(),
		BlockNumber: block.Number().Uint64(),
	}, nil
}
