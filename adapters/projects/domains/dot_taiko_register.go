package domains

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
	logMintedDomainSignature   string = "MintedDomain(string,uint256,address,uint256)"
	logNameRegisteredSignature string = "NameRegistered(uint256,string,bytes32,address,uint256,uint256)"
	logProfileCreatedSignature string = "ProfileCreated(uint256,address,address,string,string,address,bytes,string,uint256)"
)

type DotTaikoIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewDotTaikoIndexer(client *ethclient.Client) *DotTaikoIndexer {
	return &DotTaikoIndexer{
		client: client,
		addresses: []common.Address{
			// https://taikoscan.io/address/0xD7b837A0E388B4c25200983bdAa3EF3A83ca86b7
			common.HexToAddress("0xD7b837A0E388B4c25200983bdAa3EF3A83ca86b7"),
			// https://taikoscan.io/address/0xFb2Cd41a8aeC89EFBb19575C6c48d872cE97A0A5
			common.HexToAddress("0xFb2Cd41a8aeC89EFBb19575C6c48d872cE97A0A5"),
			// https://taikoscan.io/address/0x01412AAba531Cc6ef630CE5059120999f824CDAF
			common.HexToAddress("0x01412AAba531Cc6ef630CE5059120999f824CDAF"),
		}}
}

var _ adapters.LogIndexer[adapters.Whitelist] = &DotTaikoIndexer{}

func (indexer *DotTaikoIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *DotTaikoIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Whitelist, error) {
	var whitelist []adapters.Whitelist

	for _, l := range logs {
		if !indexer.isRelevantLog(l) {
			continue
		}

		owner := common.BytesToAddress(l.Topics[2].Bytes()[12:])

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		w := &adapters.Whitelist{
			User:        owner,
			Time:        block.Time(),
			BlockNumber: block.NumberU64(),
			TxHash:      l.TxHash,
		}

		whitelist = append(whitelist, *w)
	}

	return whitelist, nil
}

func (indexer *DotTaikoIndexer) isRelevantLog(l types.Log) bool {
	t := l.Topics[0].Hex()
	return t == crypto.Keccak256Hash([]byte(logMintedDomainSignature)).Hex() ||
		t == crypto.Keccak256Hash([]byte(logNameRegisteredSignature)).Hex() ||
		t == crypto.Keccak256Hash([]byte(logProfileCreatedSignature)).Hex()
}
