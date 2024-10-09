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
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/conft"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/domains"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/drips"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/gaming"
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
		return processRitsuLPIndexer(client, p.Blocknumber)
	case Izumi:
		return processIziLPIndexer(client, p.Blocknumber)
	case NewTransactionSender:
		processor := transactionsender.New(client)
		return processBlock(ctx, client, processor, p.Blocknumber)
	case NftDeployed:
		processor := nftdeployed.New(client)
		return processBlock(ctx, client, processor, p.Blocknumber)
	case GamingWhitelist:
		processor := gaming.NewGamingProcessor(client)
		return processBlock(ctx, client, processor, p.Blocknumber)
	case OrderFulfilledIndexer:
		return processOrderFulfilledIndexer(client, p.Blocknumber)
	case DotTaikoIndexer:
		indexer := domains.NewDotTaikoIndexer(client)
		return processLog(ctx, client, indexer, p.Blocknumber)
	case NewSaleIndexer:
		return processNewSaleIndexer(client, p.Blocknumber)
	case ContractDeployed:
		return processContractDeployedIndexer(client, p.Blocknumber)
	case CollectionCreated:
		return processCollectionCreatedIndexer(client, p.Blocknumber)
	case TokenSold:
		indexer := conft.NewTokenSoldIndexer(client, []common.Address{common.HexToAddress(conft.TokenSoldAddress)})
		return processLog(ctx, client, indexer, p.Blocknumber)
	case Drips:
		indexer := drips.NewLockIndexer(client, []common.Address{common.HexToAddress(drips.LockAddress)})
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

func processBlock[T any](ctx context.Context, client *ethclient.Client, processor adapters.BlockPprocessor[T], blocknumber int64) error {
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
