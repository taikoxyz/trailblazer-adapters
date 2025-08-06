package omnihub

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
	// https://taikoscan.io/address/0xb0B4B761C9e9Bf5A9194a42e944A4A6646B83919
	ContractDeployedAddress string = "0xb0B4B761C9e9Bf5A9194a42e944A4A6646B83919"

	logContractDeployedSignature string = "ContractDeployed(address)"
)

type ContractDeployedIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewContractDeployedIndexer(client *ethclient.Client, addresses []common.Address) *ContractDeployedIndexer {
	return &ContractDeployedIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.Whitelist] = &ContractDeployedIndexer{}

func (indexer *ContractDeployedIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *ContractDeployedIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Whitelist, error) {
	var whitelist []adapters.Whitelist

	for _, l := range logs {
		if !indexer.isContractDeployedLog(l) {
			continue
		}

		tx, err := indexer.client.TransactionInBlock(ctx, l.BlockHash, l.TxIndex)
		if err != nil {
			return nil, err
		}

		sender, err := indexer.client.TransactionSender(ctx, tx, l.BlockHash, l.TxIndex)
		if err != nil {
			return nil, err
		}

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
			User: sender,
		}

		whitelist = append(whitelist, *w)
	}

	return whitelist, nil
}

func (indexer *ContractDeployedIndexer) isContractDeployedLog(l types.Log) bool {
	return l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logContractDeployedSignature)).Hex()
}
