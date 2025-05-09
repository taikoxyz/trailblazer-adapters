package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/samber/lo"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	nftdeployed "github.com/taikoxyz/trailblazer-adapters/adapters/nft_deployed"
	"github.com/taikoxyz/trailblazer-adapters/adapters/pfp"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/avalon"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/conft"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/domains"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/dorahacks"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/drips"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/gaming"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/izumi"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/loopex"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/loopring"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/nfts2me"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/okx"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/omnihub"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/polaris"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/ritsu"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/robinos"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/symmetric"
	transactionsender "github.com/taikoxyz/trailblazer-adapters/adapters/transaction_sender"
)

const taikoRPC string = "https://rpc.taiko.xyz"

type prompt struct {
	Adapter     string `survey:"adapter"`
	Blocknumber int64  `survey:"blocknumber"`
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "trailblazer-adapters",
	Short: "Trailblazer Adapters CLI",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := promptUser()
		if err != nil {
			log.Fatalf("Prompt failed: %v", err)
		}

		err = executeCommand(*p)
		if err != nil {
			log.Fatalf("Execution failed: %v", err)
		}
	},
}

func init() {
	cobra.OnInitialize()
}

func promptUser() (*prompt, error) {
	qs := []*survey.Question{
		{
			Name: "adapter",
			Prompt: &survey.Select{
				Message: "Choose an adapter:",
				Options: lo.Map(adapterz(), func(a adapter, index int) string {
					return a.String()
				}),
			},
		},
		{
			Name: "blocknumber",
			Prompt: &survey.Input{
				Message: "Enter the block number:",
			},
		},
	}

	answers := prompt{}

	if err := survey.Ask(qs, &answers); err != nil {
		return nil, err
	}

	return &answers, nil
}

func executeCommand(p prompt) error {
	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	if err != nil {
		return fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	adaptr := adapter(p.Adapter)
	switch adaptr {
	case RitsuLP:
		indexer := ritsu.NewTransferIndexer(
			client,
			lo.Map(ritsu.LPAdresses(), func(a string, index int) common.Address {
				return common.HexToAddress(a)
			}),
			ritsu.Whitelist(),
		)
		return processLog(ctx, client, indexer, p.Blocknumber)

	case IzumiLP:
		indexer := izumi.NewLPTransferIndexer(
			client,
			[]common.Address{common.HexToAddress(izumi.LPAddress)},
			izumi.Whitelist(),
		)
		return processLog(ctx, client, indexer, p.Blocknumber)

	case TransactionSender:
		processor := transactionsender.New(client)
		return processBlock(ctx, client, processor, p.Blocknumber)

	case NftDeployed:
		processor := nftdeployed.New(client)
		return processBlock(ctx, client, processor, p.Blocknumber)

	case GamingWhitelist:
		processor := gaming.NewGamingProcessor(client)
		return processBlock(ctx, client, processor, p.Blocknumber)

	case OkxOrderFulfilled:
		indexer := okx.NewOrderFulfilledIndexer(
			client,
			[]common.Address{common.HexToAddress(okx.OrderFulfilledAddress)},
		)
		return processLog(ctx, client, indexer, p.Blocknumber)

	case DotTaikoDomains:
		indexer := domains.NewDotTaikoIndexer(client)
		return processLog(ctx, client, indexer, p.Blocknumber)

	case LoopexNewSale:
		indexer := loopex.NewNewSaleIndexer(
			client,
			[]common.Address{common.HexToAddress(loopex.NewSaleAddress)},
		)
		return processLog(ctx, client, indexer, p.Blocknumber)

	case OmnihubContractDeployed:
		indexer := omnihub.NewContractDeployedIndexer(
			client,
			[]common.Address{common.HexToAddress(omnihub.ContractDeployedAddress)},
		)
		return processLog(ctx, client, indexer, p.Blocknumber)

	case Nfts2meCollectionCreated:
		indexer := nfts2me.NewCollectionCreatedIndexer(
			client,
			[]common.Address{common.HexToAddress(nfts2me.CollectionCreatedAddress)},
		)
		return processLog(ctx, client, indexer, p.Blocknumber)

	case ConftTokenSold:
		indexer := conft.NewTokenSoldIndexer(
			client,
			[]common.Address{common.HexToAddress(conft.TokenSoldAddress)},
		)
		return processLog(ctx, client, indexer, p.Blocknumber)

	case DripsLock:
		indexer := drips.NewLockIndexer(
			client,
			[]common.Address{common.HexToAddress(drips.LockAddress)},
		)
		return processLog(ctx, client, indexer, p.Blocknumber)
	case SymmetricLock:
		indexer := symmetric.NewLockIndexer(
			client,
			[]common.Address{common.HexToAddress(symmetric.LockAddress)},
		)
		return processLog(ctx, client, indexer, p.Blocknumber)
	case RobinosPrediction:
		indexer := robinos.NewPredictionIndexer(
			client,
			[]common.Address{common.HexToAddress(robinos.RobinosAddress)},
			robinos.SelectedMultiplierEvents,
		)
		return processLog(ctx, client, indexer, p.Blocknumber)
	case LoopringLock:
		indexer := loopring.NewLockIndexer(
			client,
			[]common.Address{common.HexToAddress(loopring.LockAddress)},
		)
		return processLog(ctx, client, indexer, p.Blocknumber)
	case PolarisLP:
		indexer := polaris.NewLPTransferIndexer(
			client,
			[]common.Address{common.HexToAddress(polaris.VaultAddress)},
		)
		return processLog(ctx, client, indexer, p.Blocknumber)
	case DoraHacksVoting:
		indexer := dorahacks.NewVotingIndexer(
			client,
			[]common.Address{common.HexToAddress(dorahacks.VotingAddress)},
		)
		return processLog(ctx, client, indexer, p.Blocknumber)
	case AvalonClaim:
		indexer := avalon.NewClaimIndexer(
			client,
			[]common.Address{common.HexToAddress(avalon.ClaimAddress)},
		)
		return processLog(ctx, client, indexer, p.Blocknumber)
	case PfpRegister:
		indexer := pfp.NewRegisterIndexer(
			client,
			[]common.Address{common.HexToAddress(pfp.RegisterAddress)},
		)
		return processLog(ctx, client, indexer, p.Blocknumber)
	case LoopringDeposit:
		indexer := loopring.NewDepositRequestedIndexer(
			client,
			[]common.Address{common.HexToAddress(loopring.DepositRequestedAddress)},
		)
		return processLog(ctx, client, indexer, p.Blocknumber)
	default:
		return fmt.Errorf("adapter %s is not supported", p.Adapter)
	}
}

func processLog[T any](ctx context.Context, client *ethclient.Client, indexer adapters.LogIndexer[T], blocknumber int64) error {
	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	if err != nil {
		return err
	}

	processed, err := indexer.Index(ctx, logs...)
	if err != nil {
		return err
	}

	log.Printf("processed: %+v\n", processed)
	return nil
}

func processBlock[T any](ctx context.Context, client *ethclient.Client, processor adapters.BlockProcessor[T], blocknumber int64) error {
	block, err := adapters.GetBlock(ctx, client, blocknumber)
	if err != nil {
		return err
	}

	processed, err := processor.Process(ctx, block)
	if err != nil {
		return err
	}

	log.Printf("processed: %+v\n", processed)
	return nil
}
