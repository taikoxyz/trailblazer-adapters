package okidori_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/okidori"
	"github.com/taikoxyz/trailblazer-adapters/testutils"
)

func TestNftSoldIndexer(t *testing.T) {
	blocknumber := int64(1293884)

	ctx := context.Background()

	client, err := ethclient.Dial(testutils.TaikoRPC)
	require.NoError(t, err)

	indexer := okidori.NewNftSoldIndexer(client, []common.Address{common.HexToAddress(okidori.MarketplaceAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	ps, err := indexer.Index(ctx, logs...)
	require.NoError(t, err)
	require.Len(t, ps, 1)
	assert.Equal(t, common.HexToAddress("0xB66A56126fAd14563e62BA2Cda658cb97F7a90dE"), ps[0].Nft.Collection)
	assert.Equal(t, big.NewInt(99253), ps[0].Nft.TokenId)
	assert.Equal(t, common.HexToAddress("0x9aC2Af0379cC8D8075474e6487242C4dea241a91"), ps[0].Seller)
	assert.Equal(t, common.HexToAddress("0x33B8A069444890E6245A617AEEC7131c84369AB1"), ps[0].Buyer)
	assert.Equal(t, big.NewInt(200000000000000000), ps[0].Price)
	assert.Equal(t, common.HexToAddress("0xA9d23408b9bA935c230493c40C73824Df71A0975"), ps[0].Currency)
	assert.Equal(t, uint64(1753968995), ps[0].BlockTime)
	assert.Equal(t, uint64(blocknumber), ps[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0xd817133dbc199e287cf8c67b515a811911e668a8be8d2f8aeabd9c1adf7d7542"), ps[0].TxHash)
}
