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
	"Prediction - $TAIKO price on Jan 19 to be over $1.50",
	"Prediction - $BTC price on Jan 19 to be over $94,500",
	"Prediction - $ETH price on Jan 19 to be over $3,200",
	"Prediction - $TAIKO price on Jan 26 to be over $1.30",
	"Prediction - $BTC price on Jan 26 to be over $101,000",
	"Prediction - $ETH price on Jan 26 to be over $3,200",
	"UEFA Champions League 24/25 - Sparta Prague v. Inter Milan",
	"UEFA Champions League 24/25 - PSG v. Man City",
	"UEFA Champions League 24/25 - Arsenal v. Dinamo Zagreb",
	"UEFA Champions League 24/25 - AC Milan v. Girona",
	"UEFA Champions League 24/25 - Liverpool v. Lille",
	"UEFA Champions League 24/25 - Club Brugge v. Juventus",
	"UEFA Champions League 24/25 - Benfica v. Barcelona",
	"UEFA Champions League 24/25 - Monaco v. Aston Villa",
	"Prediction - $TAIKO price on Feb 2 to be over $1.30",
	"Prediction - $BTC price on Feb 2 to be over $102,000",
	"Prediction - $ETH price on Feb 2 to be over $3,200",
	"UEFA Champions League 24/25 - VfB Stuttgart v. PSG",
	"UEFA Champions League 24/25 - Man City v. Club Brugge",
	"UEFA Champions League 24/25 - PSV v. Liverpool",
	"UEFA Champions League 24/25 - Girona v. Arsenal",
	"UEFA Champions League 24/25 - Barcelona v. Atalanta",
	"UEFA Champions League 24/25 - Aston Villa v. Celtic",
	"Prediction - Will $BTC go under $91,000 before Feb 9th",
	"Prediction - Will $TAIKO be over $1.30 this week?",
	"Prediction - Will $BTC be over $100,000 this week?",
	"Prediction - Will $ETH be over $3,000 this week?",
	"UEFA Champions League 24/25 - Monaco v. Benfica",
	"UEFA Champions League 24/25 - Feyenoord v. AC Milan",
	"UEFA Champions League 24/25 - Celtic v. Bayern Munich",
	"UEFA Champions League 24/25 - Club Brugge v. Atalanta",
	"UEFA Champions League 24/25 - Sporting v. Borussia Dortmund",
	"UEFA Champions League 24/25 - Man City v. Real Madrid",
	"UEFA Champions League 24/25 - Juventus v. PSV",
	"UEFA Champions League 24/25 - Brest v. PSG",
	"Prediction - $TAIKO to go over $1.20 before Mar 10",
	"Prediction - $BTC to go over $98,000 before Mar 10",
	"Prediction - $ETH to go over $2,800 before Mar 10",
	"UEFA Champions League 24/25 - PSG v. Liverpool",
	"UEFA Champions League 24/25 - Benfica v. Barcelona 5",
	"UEFA Champions League 24/25 - Bayern Munich v. Bayer Leverkusen",
	"UEFA Champions League 24/25 - Feyenoord v. Inter Milan",
	"UEFA Champions League 24/25 - Real Madrid v. Atletico Madrid",
	"UEFA Champions League 24/25 - PSV v. Arsenal",
	"UEFA Champions League 24/25 - Borussia Dortmund v. Lille",
	"UEFA Champions League 24/25 - Club Brugge v. Aston Villa 4",
	"UEFA Champions League 24/25 - Atletico Madrid v. Real Madrid",
	"UEFA Champions League 24/25 - Aston Villa v. Club Brugge",
	"UEFA Champions League 24/25 - Arsenal v. PSV",
	"UEFA Champions League 24/25 - Lille v. Borussia Dortmund",
	"UEFA Champions League 24/25 - Liverpool v. PSG",
	"UEFA Champions League 24/25 - Inter Milan v. Feyenoord",
	"UEFA Champions League 24/25 - Bayer Leverkusen v. Bayern Munich",
	"UEFA Champions League 24/25 - Barcelona v. Benfica",
	"Prediction - $TAIKO to go under $0.50 before Mar 17",
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
				Metadata: adapters.Metadata{
					BlockTime:   block.Time(),
					BlockNumber: block.NumberU64(),
					TxHash:      l.TxHash,
				},
				User:          user,
				TokenAmount:   reward,
				TokenDecimals: adapters.TaikoTokenDecimals,
				Token:         common.HexToAddress(adapters.TaikoTokenAddress),
				EventCode:     event.EventCode,
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
