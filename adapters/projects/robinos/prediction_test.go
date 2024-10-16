package robinos_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/robinos"
)

func TestPredictionIndexer(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(442090)

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	indexer := robinos.NewPredictionIndexer(client, []common.Address{common.HexToAddress(robinos.RobinosAddress)}, robinos.SelectedMultiplierEvents)

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	predictions, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, predictions, 3)
	txHash := common.HexToHash("0xfe2cdfa28a39d23061c7bf7dffeee74dee42a58acc8c82175a6aa6ffdf24890d")
	event := "Prediction - $TAIKO price on Oct 6"
	var tokenAmount1 *big.Int
	if amount, success := new(big.Int).SetString("86617647058823529411", 10); success {
		tokenAmount1 = amount
	}
	var tokenAmount2 *big.Int
	if amount, success := new(big.Int).SetString("5197058823529411764", 10); success {
		tokenAmount2 = amount
	}
	var tokenAmount3 *big.Int
	if amount, success := new(big.Int).SetString("25985294117647058823", 10); success {
		tokenAmount3 = amount
	}
	expectedPredictions := []adapters.Prediction{
		{
			User:          common.HexToAddress("0x7e00AcD60C9d9D3339B53ec9d4b691df81ae32bf"),
			TokenAmount:   tokenAmount1,
			TokenDecimals: adapters.TaikoTokenDecimals,
			Token:         common.HexToAddress(adapters.TaikoTokenAddress),
			EventCode:     event,
			BlockTime:     uint64(1728292775),
			BlockNumber:   uint64(blocknumber),
			TxHash:        txHash,
		},
		{
			User:          common.HexToAddress("0xB0A4Ce05285ED7614AA88BA9F1C1741b34971a15"),
			TokenAmount:   tokenAmount2,
			TokenDecimals: adapters.TaikoTokenDecimals,
			Token:         common.HexToAddress(adapters.TaikoTokenAddress),
			EventCode:     event,
			BlockTime:     uint64(1728292775),
			BlockNumber:   uint64(blocknumber),
			TxHash:        txHash,
		},
		{
			User:          common.HexToAddress("0x3255f62BbD1317ed1590228f287be8D602d78443"),
			TokenAmount:   tokenAmount3,
			TokenDecimals: adapters.TaikoTokenDecimals,
			Token:         common.HexToAddress(adapters.TaikoTokenAddress),
			EventCode:     event,
			BlockTime:     uint64(1728292775),
			BlockNumber:   uint64(blocknumber),
			TxHash:        txHash,
		},
	}

	for i, prediction := range predictions {
		assert.Equal(t, expectedPredictions[i].User, prediction.User)
		assert.Equal(t, expectedPredictions[i].TokenAmount, prediction.TokenAmount)
		assert.Equal(t, expectedPredictions[i].TokenDecimals, prediction.TokenDecimals)
		assert.Equal(t, expectedPredictions[i].Token, prediction.Token)
		assert.Equal(t, expectedPredictions[i].EventCode, prediction.EventCode)
		assert.Equal(t, expectedPredictions[i].BlockTime, prediction.BlockTime)
		assert.Equal(t, expectedPredictions[i].BlockNumber, prediction.BlockNumber)
		assert.Equal(t, expectedPredictions[i].TxHash, prediction.TxHash)
	}
}
