package erc20transfer_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	erc20transfer "github.com/taikoxyz/trailblazer-adapters/adapters/erc20_transfer"
	"github.com/taikoxyz/trailblazer-adapters/testutils"
)

func TestIndexer(t *testing.T) {
	blocknumber := int64(447659)

	ctx := context.Background()

	client, err := ethclient.Dial(testutils.TaikoRPC)
	require.NoError(t, err)

	indexer := erc20transfer.New(client)

	logs, err := adapters.GetLogs(ctx, client, []common.Address{common.HexToAddress(adapters.TaikoTokenAddress)}, blocknumber)
	require.NoError(t, err)

	whitelist, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, whitelist, 3)
}
