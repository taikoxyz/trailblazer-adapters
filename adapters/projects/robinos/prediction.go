package robinos

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/robinos"
)

const (
	RobinosAddress                = "0x9EF9c57754eD079D750016b802DcCD45d0AB66f8"
	LogRewardDistributedSignature = "RewardDistributed(string,uint256[],address[])"
)

var SelectedMultiplierEvents = []string{
	"Prediction - $TAIKO price on Oct 6",
	"Prediction - $TAIKO price on Oct 20",
	"UEFA Champions League 24/25 - RB Leipzig v. Liverpool",
	"UEFA Champions League 24/25 - Barcelona v. Bayern Munich",
	"UEFA Champions League 24/25 - Real Madrid v. Borussia Dortmund",
	"UEFA Champions League 24/25 - PSG v. PSV",
	"UEFA Champions League 24/25 - Juventus v. VfB Stuttgart",
	"UEFA Champions League 24/25 - Aston Villa v. Bologna",
	"UEFA Champions League 24/25 - AC Milan v. Club Brugge",
	"Prediction - $TAIKO price on Oct 27",
	"USA Elections 2024 - Trump v. Harris",
	"Prediction - $TAIKO price on Nov 3",
	"Prediction - $TAIKO price on Nov 10",
	"UEFA Champions League 24/25 - Lille v. Juventus",
	"UEFA Champions League 24/25 - Bayern v. Benfica",
	"UEFA Champions League 24/25 - PSG v. Atletico Madrid",
	"UEFA Champions League 24/25 - Inter Milan v. Arsenal",
	"UEFA Champions League 24/25 - Club Brugge v. Aston Villa",
	"UEFA Champions League 24/25 - Sporting v. Man City",
	"UEFA Champions League 24/25 - Real Madrid v. AC Milan",
	"UEFA Champions League 24/25 - Liverpool v. Bayer Leverkusen",
	"Prediction - $TAIKO price on Nov 17 B",
	"Prediction - $TAIKO price on Nov 17 A",
	"Prediction - $BTC price on Nov 24",
	"Prediction - $ETH price on Nov 24",
	"Prediction - $TAIKO price on Nov 24",
	"Prediction - $BTC price on Dec 8",
	"Prediction - $ETH price on Dec 8",
	"Prediction - $TAIKO price on Dec 8",
	"UEFA Champions League 24/25 - Juventus v. Man City",
	"UEFA Champions League 24/25 - Borussia Dortmund v. Barcelona",
	"UEFA Champions League 24/25 - RB Leipzig v. Aston Villa",
	"UEFA Champions League 24/25 - Girona v. Liverpool",
	"Prediction - $BTC price on Dec 15",
	"Prediction - $ETH price on Dec 15",
	"Prediction - $TAIKO price on Dec 15",
	"Prediction - $TAIKO price on Dec 22",
	"Prediction - $BTC price on Dec 22",
	"Prediction - $ETH price on Dec 22",
	"Prediction - $TAIKO price on Jan 12 to be over $1.75",
	"Prediction - $BTC price on Jan 12 to be over $98,500",
	"Prediction - $ETH price on Jan 12 to be over $3,600",
}

type PredictionIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
	events    []string
}

func NewPredictionIndexer(client *ethclient.Client, addresses []common.Address, events []string) *PredictionIndexer {
	return &PredictionIndexer{
		client:    client,
		addresses: addresses,
		events:    events,
	}
}

var _ adapters.LogIndexer[adapters.Prediction] = &PredictionIndexer{}

func (indexer *PredictionIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *PredictionIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Prediction, error) {
	var predictions []adapters.Prediction
	for _, l := range logs {
		if !indexer.isRewardDistributedEvent(l) {
			continue
		}

		// Unpack the reward distributed event
		var event struct {
			EventCode    string
			UserRewards  []*big.Int
			WinningUsers []common.Address
		}
		if err := processRewardDistributedLog(l, &event); err != nil {
			return nil, err
		}

		if !indexer.isSelectedMultiplierEvent(event.EventCode) {
			continue
		}

		// Fetch block details
		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		// Process each user and their corresponding reward
		for i, user := range event.WinningUsers {
			reward := event.UserRewards[i]

			// Create a Prediction struct
			prediction := adapters.Prediction{
				User:          user,
				TokenAmount:   reward,
				TokenDecimals: adapters.TaikoTokenDecimals,
				Token:         common.HexToAddress(adapters.TaikoTokenAddress),
				EventCode:     event.EventCode,
				BlockTime:     block.Time(),
				BlockNumber:   block.Number().Uint64(),
				TxHash:        l.TxHash,
			}

			predictions = append(predictions, prediction)
		}
	}

	return predictions, nil
}

// Helper function to unpack the RewardDistributed event from the log
func processRewardDistributedLog(vLog types.Log, event *struct {
	EventCode    string
	UserRewards  []*big.Int
	WinningUsers []common.Address
}) error {
	contractABI, err := abi.JSON(strings.NewReader(robinos.RobinosABI))
	if err != nil {
		return err
	}
	return contractABI.UnpackIntoInterface(event, "RewardDistributed", vLog.Data)
}

// Check if the log is a RewardDistributed event
func (indexer *PredictionIndexer) isRewardDistributedEvent(l types.Log) bool {
	return len(l.Topics) > 0 && l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(LogRewardDistributedSignature)).Hex()
}

func (indexer *PredictionIndexer) isSelectedMultiplierEvent(eventCode string) bool {
	for _, selectedEvent := range indexer.events {
		if selectedEvent == eventCode {
			return true
		}
	}
	return false
}
