package pfp_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/pfp"
	"github.com/taikoxyz/trailblazer-adapters/testutils"
)

func TestRegisterIndexer(t *testing.T) {
	blocknumber := int64(987097)

	ctx := context.Background()

	client, err := ethclient.Dial(testutils.TaikoRPC)
	require.NoError(t, err)

	indexer := pfp.NewRegisterIndexer(client, []common.Address{common.HexToAddress(pfp.RegisterAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	ws, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, ws, 1)
	assert.Equal(t, common.HexToAddress("0x1b770D723AfaB2AC34E62BBCfF5eB432BA81f18b"), ws[0].User)
	assert.Equal(t, uint64(987097), ws[0].BlockNumber)
	assert.Equal(t, uint64(1742208011), ws[0].BlockTime)
	assert.Equal(t, common.HexToHash("0xd050f45de8e6b62516a1dd77d1b43a1ab12fe5a266c365ba857b66730c2e7194"), ws[0].TxHash)
}
