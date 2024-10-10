package gaming_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/gaming"
)

func TestGamingProcessor(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(440972)

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	processor := gaming.NewGamingProcessor(client)

	block, err := adapters.GetBlock(ctx, client, blocknumber)
	require.NoError(t, err)

	whitelist, err := processor.Process(ctx, block)
	assert.NoError(t, err)
	assert.Len(t, whitelist, 25)
}
