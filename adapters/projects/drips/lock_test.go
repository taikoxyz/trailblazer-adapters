package drips_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/drips"
)

func TestLockIndexer(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(445053)

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	indexer := drips.NewLockIndexer(client, []common.Address{common.HexToAddress(drips.LockAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	locks, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, locks, 1)
	assert.Equal(t, common.HexToAddress("0xC3204E92B0e7731d75Ad667a93c8Da815BD9Ac61"), locks[0].User)
	assert.Equal(t, big.NewInt(2000000000000000000), locks[0].TokenAmount)
	assert.Equal(t, adapters.TaikoTokenDecimals, locks[0].TokenDecimals)
	assert.Equal(t, common.HexToAddress(adapters.TaikoTokenAddress), locks[0].Token)
	assert.Equal(t, uint64(1728390191), locks[0].Time)
	assert.Equal(t, uint64(blocknumber), locks[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0x95f528b52f0a75176543f516014bbba26e003f1c17c9b9413e936240e3f44650"), locks[0].TxHash)
}
