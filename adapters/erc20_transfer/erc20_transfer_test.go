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
)

func TestIndexer(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(447659)

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	indexer := erc20transfer.New(client)

	logs, err := adapters.GetLogs(ctx, client, []common.Address{common.HexToAddress(adapters.TaikoTokenAddress)}, blocknumber)
	require.NoError(t, err)

	whitelist, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, whitelist, 3)
}
