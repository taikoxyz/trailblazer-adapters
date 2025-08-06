package nfts2me_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/nfts2me"
	"github.com/taikoxyz/trailblazer-adapters/testutils"
)

func TestCollectionCreatedIndexer(t *testing.T) {
	blocknumber := int64(447218)

	ctx := context.Background()

	client, err := ethclient.Dial(testutils.TaikoRPC)
	require.NoError(t, err)

	indexer := nfts2me.NewCollectionCreatedIndexer(client, []common.Address{common.HexToAddress(nfts2me.CollectionCreatedAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	whitelist, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, whitelist, 1)
	assert.Equal(t, common.HexToAddress("0x49ea72fd31ecfa8e3cC3FEC8827BBCD9AEbc078B"), whitelist[0].User)
	assert.Equal(t, uint64(1728462755), whitelist[0].BlockTime)
	assert.Equal(t, uint64(blocknumber), whitelist[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0x753bcce887f5da1abedbc3f620f558586ad05ed5009a6e77c1722493e9d47c76"), whitelist[0].TxHash)
}
