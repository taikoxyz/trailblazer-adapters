package domains_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/domains"
)

func TestDotTaikoIndexer_NameRegistered(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(437536)

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	indexer := domains.NewDotTaikoIndexer(client)

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	whitelist, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, whitelist, 1)
	assert.Equal(t, common.HexToAddress("0x4C2d0F9080b4D3299EF51F8684641024522f4549"), whitelist[0].User)
	assert.Equal(t, uint64(1728148079), whitelist[0].Time)
	assert.Equal(t, uint64(blocknumber), whitelist[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0x3790b1e22c54773db2e6ab4d11a026b07e861a2e509faca24a6789379629d5bf"), whitelist[0].TxHash)
}

func TestDotTaikoIndexer_MintedDomain(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(447301)

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	indexer := domains.NewDotTaikoIndexer(client)

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	whitelist, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, whitelist, 1)
	assert.Equal(t, common.HexToAddress("0x5A27d268e830655E908a0a2C3B24F572695AF5e8"), whitelist[0].User)
	assert.Equal(t, uint64(1728465575), whitelist[0].Time)
	assert.Equal(t, uint64(blocknumber), whitelist[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0x52d0dbf1052e413148044eee1322f288d9338e81a0af0edd9567da02e3922f4a"), whitelist[0].TxHash)
}

func TestDotTaikoIndexer_ProfileCreated(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(447625)

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	indexer := domains.NewDotTaikoIndexer(client)

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	whitelist, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, whitelist, 1)
	assert.Equal(t, common.HexToAddress("0x25a0E0971bB44cd668BA27e974982dDFDB7f916D"), whitelist[0].User)
	assert.Equal(t, uint64(1728476387), whitelist[0].Time)
	assert.Equal(t, uint64(blocknumber), whitelist[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0xd8d1b12cc44061097ded9e19032a1e7ae4a691ea00e49e7dbe02af9457340734"), whitelist[0].TxHash)
}
