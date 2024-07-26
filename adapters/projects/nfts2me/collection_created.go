package nfts2me

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
)

var (
	logTransferSigHash = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
)

type CollectionCreatedIndexer struct {
	TargetAddresses []common.Address
}

// NewCollectionCreatedIndexer creates a new CollectionCreatedIndexer.
func NewCollectionCreatedIndexer() *CollectionCreatedIndexer {
	return &CollectionCreatedIndexer{TargetAddresses: []common.Address{common.HexToAddress("0x00000000001594C61dD8a6804da9AB58eD2483ce")}}
}

func (indexer *CollectionCreatedIndexer) Addresses() []common.Address {
	return indexer.TargetAddresses
}

func (indexer *CollectionCreatedIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.Whitelist, error) {
	var result []adapters.Whitelist
	for _, vLog := range logs {
		if !isERC721Transfer(vLog) && !isFromZeroAddress(vLog) {
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

func isERC721Transfer(vLog types.Log) bool {
	return len(vLog.Topics) == 4 && vLog.Topics[0].Hex() == logTransferSigHash.Hex()
}

func isFromZeroAddress(vLog types.Log) bool {
	from := common.BytesToAddress(vLog.Topics[1].Bytes()[12:])
	return from.Hex() != adapters.ZeroAddress.Hex()
}

// processLog processes a single ERC20 transfer log.
func (indexer *CollectionCreatedIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.Whitelist, error) {
	to := common.BytesToAddress(vLog.Topics[2].Bytes()[12:])

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	return &adapters.Whitelist{
		User:        to,
		Time:        block.Time(),
		BlockNumber: block.Number().Uint64(),
	}, nil
}
