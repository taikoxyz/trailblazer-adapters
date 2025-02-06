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
	// TODO: update with correct addresses and decimal
	ClaimAddress string = "0x46f0a2e45bee8e9ebfdb278ce06caa6af294c349"
	// L1: https://etherscan.io/token/0x5c8d0c48810fd37a0a824d074ee290e64f7a8fa2
	// L2: https://taikoscan.io/token/0xE9cA67e5051e1806546d0a06ee465221c5877feE
	AvlTokenAddress string = "0xE9cA67e5051e1806546d0a06ee465221c5877feE"
	AvlTokenDecimal uint8  = 18

	logClaimedSignature string = "Claimed(address,uint256,uint256)"
)

type ClaimIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewClaimIndexer(client *ethclient.Client, addresses []common.Address) *ClaimIndexer {
	return &ClaimIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.Position] = &ClaimIndexer{}

func (indexer *ClaimIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *ClaimIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Position, error) {
	var claimedEvent struct {
		AvlAmount  *big.Int
		UsdaAmount *big.Int
	}

	var claims []adapters.Position

	for _, l := range logs {
		if !indexer.isClaimed(l) {
			continue
		}

		user := common.BytesToAddress(l.Topics[1].Bytes()[12:])

		erc20ABI, err := abi.JSON(strings.NewReader(erc20.ABI))
		if err != nil {
			return nil, err
		}

		err = erc20ABI.UnpackIntoInterface(&claimedEvent, "Claimed", l.Data)
		if err != nil {
			return nil, err
		}

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		claim := &adapters.Position{
			User:          user,
			TokenAmount:   claimedEvent.AvlAmount,
			TokenDecimals: AvlTokenDecimal,
			Token:         common.HexToAddress(AvlTokenAddress),
			BlockTime:     block.Time(),
			BlockNumber:   block.NumberU64(),
			TxHash:        l.TxHash,
		}

		claims = append(claims, *claim)
	}

	return claims, nil
}

func (indexer *ClaimIndexer) isClaimed(l types.Log) bool {
	return l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logClaimedSignature)).Hex()
}
