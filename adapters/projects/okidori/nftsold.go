package okidori

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/okidori_marketplace"
)

const (
	// MarketplaceAddress is the address of the staging environment Okidori NFT marketplace on Taiko Alethia.
	// https://taikoscan.io/address/0xa98221F14BDf98cAeD75Fdc99778c503b9398109
	MarketplaceAddress string = "0xa98221F14BDf98cAeD75Fdc99778c503b9398109"
)

type NftSoldIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewNftSoldIndexer(client *ethclient.Client, addresses []common.Address) *NftSoldIndexer {
	return &NftSoldIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.NftSold] = &NftSoldIndexer{}

func (indexer *NftSoldIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *NftSoldIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.NftSold, error) {
	var nftSolds []adapters.NftSold

	okidoriMarketplaceABI, err := abi.JSON(strings.NewReader(okidori_marketplace.OkidoriMarketplaceABI))
	if err != nil {
		return nil, err
	}
	nftSoldEvent := okidoriMarketplaceABI.Events["NFTSold"]
	nftSoldEventSigHash := nftSoldEvent.ID

	for _, l := range logs {
		if l.Topics[0].Hex() != nftSoldEventSigHash.Hex() {
			continue
		}

		// Topics
		listingId := new(big.Int).SetBytes(l.Topics[1].Bytes())
		_ = listingId // not used, but kept for completeness
		buyer := common.HexToAddress(l.Topics[2].Hex())
		seller := common.HexToAddress(l.Topics[3].Hex())

		// Data
		var nftSoldEventData struct {
			ContractAddress common.Address
			TokenId         *big.Int
			Price           *big.Int
			Currency        common.Address
		}

		err = okidoriMarketplaceABI.UnpackIntoInterface(&nftSoldEventData, "NFTSold", l.Data)
		if err != nil {
			return nil, err
		}

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		nftSold := &adapters.NftSold{
			Metadata: adapters.Metadata{
				BlockTime:   block.Time(),
				BlockNumber: block.NumberU64(),
				TxHash:      l.TxHash,
			},
			Nft: adapters.Nft{
				Collection: nftSoldEventData.ContractAddress,
				TokenId:    nftSoldEventData.TokenId,
			},
			Buyer:    buyer,
			Seller:   seller,
			Price:    nftSoldEventData.Price,
			Currency: nftSoldEventData.Currency,
		}

		nftSolds = append(nftSolds, *nftSold)
	}

	return nftSolds, nil
}
