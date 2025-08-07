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

const (
	// https://taikoscan.io/address/0x00000000001594C61dD8a6804da9AB58eD2483ce
	CollectionCreatedAddress = "0x00000000001594C61dD8a6804da9AB58eD2483ce"

	logTransferSignature string = "Transfer(address,address,uint256)"
)

type CollectionCreatedIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewCollectionCreatedIndexer(client *ethclient.Client, addresses []common.Address) *CollectionCreatedIndexer {
	return &CollectionCreatedIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.Whitelist] = &CollectionCreatedIndexer{}

func (indexer *CollectionCreatedIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *CollectionCreatedIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Whitelist, error) {
	var whitelist []adapters.Whitelist

	for _, l := range logs {
		if !indexer.isERC721Transfer(l) && !indexer.isFromZeroAddress(l) {
			continue
		}

		to := common.BytesToAddress(l.Topics[2].Bytes()[12:])

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		w := &adapters.Whitelist{
			Metadata: adapters.Metadata{
				BlockTime:   block.Time(),
				BlockNumber: block.NumberU64(),
				TxHash:      l.TxHash,
			},
			User: to,
		}

		whitelist = append(whitelist, *w)
	}

	return whitelist, nil
}

func (indexer *CollectionCreatedIndexer) isERC721Transfer(l types.Log) bool {
	return len(l.Topics) == 4 && l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logTransferSignature)).Hex()
}

func (indexer *CollectionCreatedIndexer) isFromZeroAddress(l types.Log) bool {
	from := common.BytesToAddress(l.Topics[1].Bytes()[12:])
	return from.Hex() != adapters.ZeroAddress().Hex()
}
