package omnihub_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/omnihub"
)

func TestTokenSoldIndexer(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(447165)

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	indexer := omnihub.NewContractDeployedIndexer(client, []common.Address{common.HexToAddress(omnihub.ContractDeployedAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	whitelist, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, whitelist, 1)
	assert.Equal(t, common.HexToAddress("0x8bC4787A5FE09f8ACa022404BEF7021fDa167593"), whitelist[0].User)
	assert.Equal(t, uint64(1728460811), whitelist[0].Time)
	assert.Equal(t, uint64(blocknumber), whitelist[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0x2db96a2cab82e3ff7bd9574563268d5ed7615adfe4ba2693885644a5d650486c"), whitelist[0].TxHash)
}
