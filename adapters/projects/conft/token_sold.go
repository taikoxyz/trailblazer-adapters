package conft

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
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/conft"
)

var (
	logTokenSoldSigHash = crypto.Keccak256Hash([]byte("TokenSold(uint128,address,address,address,uint256,uint128)"))
)

type TokenSoldEvent struct {
	ID              *big.Int
	Seller          common.Address
	Buyer           common.Address
	ContractAddress common.Address
	TokenID         *big.Int
	Price           *big.Int
}

type TokenSoldIndexer struct {
	TargetAddresses []common.Address
}

// NewTokenSoldIndexer creates a new TokenSoldIndexer.
func NewTokenSoldIndexer() *TokenSoldIndexer {
	return &TokenSoldIndexer{TargetAddresses: []common.Address{common.HexToAddress("0x6Ce2CFD7674cf47A851690a11d1DB45c6cCBe17A")}}
}

func (indexer *TokenSoldIndexer) Addresses() []common.Address {
	return indexer.TargetAddresses
}

func (indexer *TokenSoldIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.Whitelist, error) {
	var result []adapters.Whitelist
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

func (indexer *TokenSoldIndexer) isRelevantLog(topic common.Hash) bool {
	return topic.Hex() == logTokenSoldSigHash.Hex()
}

func (indexer *TokenSoldIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.Whitelist, error) {
	var tokenSoldEvent TokenSoldEvent

	tokenSoldABI, err := abi.JSON(strings.NewReader(conft.ABI))
	if err != nil {
		return nil, err
	}

	err = tokenSoldABI.UnpackIntoInterface(&tokenSoldEvent, "OrderFulfilled", vLog.Data)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	return &adapters.Whitelist{
		User:        tokenSoldEvent.Buyer,
		Time:        block.Time(),
		BlockNumber: block.Number().Uint64(),
	}, nil
}
