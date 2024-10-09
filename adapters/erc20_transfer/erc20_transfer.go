package erc20transfer

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
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/erc20"
)

const logTransferSignature string = "Transfer(address,address,uint256)"

type Indexer struct {
	client *ethclient.Client
}

func New(client *ethclient.Client) *Indexer {
	return &Indexer{
		client: client,
	}
}

var _ adapters.LogIndexer[adapters.Whitelist] = &Indexer{}

func (indexer *Indexer) Addresses() []common.Address {
	return nil
}

func (indexer *Indexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Whitelist, error) {
	var whitelist []adapters.Whitelist

	for _, l := range logs {
		if !isERC20Transfer(l) {
			continue
		}

		var transferEvent struct {
			Value *big.Int
		}

		to := common.BytesToAddress(l.Topics[2].Bytes()[12:])

		erc20ABI, err := abi.JSON(strings.NewReader(erc20.ABI))
		if err != nil {
			return nil, err
		}

		err = erc20ABI.UnpackIntoInterface(&transferEvent, "Transfer", l.Data)
		if err != nil {
			return nil, err
		}

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		w := &adapters.Whitelist{
			User:        to,
			Time:        block.Time(),
			BlockNumber: block.NumberU64(),
			TxHash:      l.TxHash,
		}

		whitelist = append(whitelist, *w)
	}

	return whitelist, nil
}

func isERC20Transfer(l types.Log) bool {
	return len(l.Topics) == 3 && l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logTransferSignature)).Hex()
}
