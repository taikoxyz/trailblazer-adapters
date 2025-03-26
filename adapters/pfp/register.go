package pfp

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
)

const (
	// https://taikoscan.io/address/0x58617427f3d42e5435908661d3c788d7d2EAf3fa
	RegisterAddress string = "0x58617427f3d42e5435908661d3c788d7d2EAf3fa"

	// event ProfilePictureSet(address indexed user, address indexed nftContract, uint256 indexed tokenId);
	logProfilePictureSetSignatureHex string = "0x69cad6da4174faa68461906571bf7b6151f1b2958fe0c1e1b0ca11ad228d270b"
)

type RegisterIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewRegisterIndexer(client *ethclient.Client, addresses []common.Address) *RegisterIndexer {
	return &RegisterIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.Whitelist] = &RegisterIndexer{}

func (indexer *RegisterIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *RegisterIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Whitelist, error) {
	var ws []adapters.Whitelist

	for _, l := range logs {
		if !indexer.isProfilePictureSet(l) {
			continue
		}

		user := common.BytesToAddress(l.Topics[1].Bytes()[12:])

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		w := &adapters.Whitelist{
			User:        user,
			Time:        block.Time(),
			BlockNumber: block.NumberU64(),
			TxHash:      l.TxHash,
		}

		ws = append(ws, *w)
	}

	return ws, nil
}

func (indexer *RegisterIndexer) isProfilePictureSet(l types.Log) bool {
	return l.Topics[0].Hex() == logProfilePictureSetSignatureHex
}
