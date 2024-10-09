package loopex_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/loopex"
)

func TestNewSaleIndexer(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(442580)

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	indexer := loopex.NewNewSaleIndexer(client, []common.Address{common.HexToAddress(loopex.NewSaleAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	whitelist, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, whitelist, 1)
	assert.Equal(t, common.HexToAddress("0x83574780fcb69dD5735F7Aa37005fF8E53b866f9"), whitelist[0].User)
	assert.Equal(t, uint64(1728308555), whitelist[0].Time)
	assert.Equal(t, uint64(blocknumber), whitelist[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0x3453db756608c261722e5cf8e8d0df5e9b934dac45304c67f5834b96c87134bf"), whitelist[0].TxHash)
}
