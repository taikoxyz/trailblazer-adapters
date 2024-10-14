package nomis_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/nomis"
)

func TestTokenSoldIndexer(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(448072)

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	indexer := nomis.NewScoreMintedIndexer(client, []common.Address{common.HexToAddress(nomis.ScoreMintedAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	whitelist, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, whitelist, 8)
}