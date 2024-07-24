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
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/sale"
)

var (
	logNewSaleSigHash = crypto.Keccak256Hash([]byte("NewSale(address,uint256,address,uint256,address,uint256,uint256)"))
)

type NewSaleIndexer struct {
	TargetAddresses []common.Address
}

// NewSaleEvent represents the NewSale event structure
type NewSaleEvent struct {
	Seller    common.Address `json:"seller"`
	TokenId   *big.Int       `json:"tokenId"`
	Buyer     common.Address `json:"buyer"`
	Price     *big.Int       `json:"price"`
	Currency  common.Address `json:"currency"`
	Amount    *big.Int       `json:"amount"`
	Timestamp *big.Int       `json:"timestamp"`
}

func NewNewSaleIndexer() *NewSaleIndexer {
	return &NewSaleIndexer{TargetAddresses: []common.Address{
		common.HexToAddress("0x8733764434c3a3C2b2Fa3A5033CA49FDdDF976C0"),
	}}
}

func (indexer *NewSaleIndexer) Addresses() []common.Address {
	return indexer.TargetAddresses
}

func (indexer *NewSaleIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.TransferData, error) {
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

func (indexer *NewSaleIndexer) isRelevantLog(topic common.Hash) bool {
	return topic.Hex() == logNewSaleSigHash.Hex()
}

func (indexer *NewSaleIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.TransferData, error) {
	var newSaleEvent NewSaleEvent

	newSaleABI, err := abi.JSON(strings.NewReader(sale.ABI))
	if err != nil {
		return nil, err
	}

	err = newSaleABI.UnpackIntoInterface(&newSaleEvent, "NewSale", vLog.Data)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	return &adapters.TransferData{
		To:          newSaleEvent.Buyer,
		Time:        block.Time(),
		BlockNumber: block.Number().Uint64(),
	}, nil
}
