package avalon

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

const (
	// TODO: update
	AvalonAirdropContractAddress string = "0x46f0a2e45bee8e9ebfdb278ce06caa6af294c349"
	AvalonTokenAddress           string = "0x46f0a2e45bee8e9ebfdb278ce06caa6af294c349"
	AvalonTokenDecimal                  = 18

	logTransferSignature string = "Transfer(address,address,uint256)"
)

type ClaimIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
	contract  common.Address
}

func NewClaimIndexer(client *ethclient.Client, contract common.Address, addresses []common.Address) *ClaimIndexer {
	return &ClaimIndexer{
		client:    client,
		addresses: addresses,
		contract:  contract,
	}
}

var _ adapters.LogIndexer[adapters.Position] = &ClaimIndexer{}

func (indexer *ClaimIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *ClaimIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Position, error) {
	var claims []adapters.Position

	for _, l := range logs {
		if !indexer.isTransfer(l) {
			continue
		}

		var transferEvent struct {
			Value *big.Int
		}

		to := common.BytesToAddress(l.Topics[2].Bytes()[12:])
		from := common.BytesToAddress(l.Topics[1].Bytes()[12:])

		if from != indexer.contract {
			continue
		}

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

		claim := &adapters.Position{
			User:          to,
			TokenAmount:   transferEvent.Value,
			TokenDecimals: AvalonTokenDecimal,
			Token:         common.HexToAddress(adapters.TaikoTokenAddress),
			BlockTime:     block.Time(),
			BlockNumber:   block.NumberU64(),
			TxHash:        l.TxHash,
		}

		claims = append(claims, *claim)
	}

	return claims, nil
}

func (indexer *ClaimIndexer) isTransfer(l types.Log) bool {
	return l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logTransferSignature)).Hex()
}
