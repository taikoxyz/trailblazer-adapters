package conft_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/conft"
	"github.com/taikoxyz/trailblazer-adapters/testutils"
)

func TestTokenSoldIndexer(t *testing.T) {
	blocknumber := int64(273007)

	ctx := context.Background()

	client, err := ethclient.Dial(testutils.TaikoRPC)
	require.NoError(t, err)

	indexer := conft.NewTokenSoldIndexer(client, []common.Address{common.HexToAddress(conft.TokenSoldAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	whitelist, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, whitelist, 1)
	assert.Equal(t, common.HexToAddress("0x1A2929B1Da32A2DcCE6442a90853De48006935dd"), whitelist[0].User)
	assert.Equal(t, uint64(1723357103), whitelist[0].BlockTime)
	assert.Equal(t, uint64(blocknumber), whitelist[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0xf66d70765a1e8dd7035678a7fe5a04fbe5cd6bbe3245d91436fa1fd46fd65cd4"), whitelist[0].TxHash)
}
