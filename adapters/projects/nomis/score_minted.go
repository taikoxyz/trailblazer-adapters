package nomis

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
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/order"
)

const (
	// https://taikoscan.io/address/0x128B7290CdFA241E77337e0e42100bf6237a783C
	ScoreMintedAddress string = "0x128B7290CdFA241E77337e0e42100bf6237a783C"

	// logScoreMintedSignature string = "ScoreMinted(bytes32,address,address,address,(uint8,address,uint256,uint256)[],(uint8,address,uint256,uint256,address)[])"
	logScoreMintedSignature string = "ChangedScore(uint256, address, uint16, uint16, uint256, string, string, string)"
)

type ChangedScoreEvent struct {
	TokenId             *big.Int
	Owner               common.Address
	Score               uint16
	CalculationModel    uint16
	ChainId             *big.Int
	MetadataUrl         string
	ReferralCode        string
	ReferrerCode        string
}

type ScoreMintedIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewScoreMintedIndexer(client *ethclient.Client, addresses []common.Address) *ScoreMintedIndexer {
	return &ScoreMintedIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.Whitelist] = &ScoreMintedIndexer{}

func (indexer *ScoreMintedIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *ScoreMintedIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Whitelist, error) {
	var whitelist []adapters.Whitelist

	for _, l := range logs {
		if !indexer.isScoreMintedLog(l) {
			continue
		}

		var changedScoreEvent ChangedScoreEvent

		scoreMintedABI, err := abi.JSON(strings.NewReader(order.ABI))
		if err != nil {
			return nil, err
		}
		err = scoreMintedABI.UnpackIntoInterface(&changedScoreEvent, "ScoreMinted", l.Data)
		if err != nil {
			return nil, err
		}

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		w := &adapters.Whitelist{
			User:        changedScoreEvent.Owner,
			Time:        block.Time(),
			BlockNumber: block.NumberU64(),
			TxHash:      l.TxHash,
		}

		whitelist = append(whitelist, *w)
	}

	return whitelist, nil
}

func (indexer *ScoreMintedIndexer) isScoreMintedLog(l types.Log) bool {
	return l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logScoreMintedSignature)).Hex()
}