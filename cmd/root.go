package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/taikoxyz/trailblazer-adapters/adapters/blocks"
	"github.com/taikoxyz/trailblazer-adapters/adapters/logs"
)

var adapter string
var blockNumber int64

var rootCmd = &cobra.Command{
	Use:   "trailblazer-adapters",
	Short: "Trailblazer Adapters CLI",
	Run: func(cmd *cobra.Command, args []string) {
		promptUser()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize()
}

func promptUser() {
	var adapterOptions = []string{"NewTransactionSender", "DotTaikoIndexer", "OrderFulfilledIndexer"}
	var qs = []*survey.Question{
		{
			Name: "adapter",
			Prompt: &survey.Select{
				Message: "Choose an adapter:",
				Options: adapterOptions,
			},
		},
		{
			Name:   "blockNumber",
			Prompt: &survey.Input{Message: "Enter the block number:"},
		},
	}

	answers := struct {
		Adapter     string `survey:"adapter"`
		BlockNumber int64  `survey:"blockNumber"`
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		log.Fatalf("Prompt failed %v", err)
	}

	adapter = answers.Adapter
	blockNumber = answers.BlockNumber

	executeCommand()
}

func executeCommand() {
	client, err := ethclient.Dial("https://rpc.taiko.xyz")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	switch adapter {
	case "NewTransactionSender":
		processor := blocks.NewTransactionSender()
		blockNumberBig := big.NewInt(blockNumber)
		block, err := client.BlockByNumber(context.Background(), blockNumberBig)
		if err != nil {
			log.Fatalf("Failed to fetch the block: %v", err)
		}

		senders, err := processor.ProcessBlock(context.Background(), block, client)
		if err != nil {
			log.Fatalf("Failed to process the block: %v", err)
		}

		fmt.Printf("Senders: %v\n", senders)
	case "OrderFulfilledIndexer":
		processor := logs.NewOrderFulfilledIndexer()

		chainID, err := client.ChainID(context.Background())
		if err != nil {
			log.Fatalf("Failed to fetch the chain ID: %v", err)
		}
		query := ethereum.FilterQuery{
			Addresses: processor.Addresses,
			FromBlock: big.NewInt(blockNumber),
			ToBlock:   big.NewInt(blockNumber),
		}
		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatalf("Failed to fetch the logs: %v", err)
		}
		senders, err := processor.IndexLogs(context.Background(), chainID, client, logs)
		if err != nil {
			log.Fatalf("Failed to process the logs: %v", err)
		}

		fmt.Printf("Senders: %v\n", senders)
	case "DotTaikoIndexer":
		processor := logs.NewDotTaikoIndexer()

		chainID, err := client.ChainID(context.Background())
		if err != nil {
			log.Fatalf("Failed to fetch the chain ID: %v", err)
		}
		query := ethereum.FilterQuery{
			Addresses: processor.Addresses,
			FromBlock: big.NewInt(blockNumber),
			ToBlock:   big.NewInt(blockNumber),
		}
		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatalf("Failed to fetch the logs: %v", err)
		}
		senders, err := processor.IndexLogs(context.Background(), chainID, client, logs)
		if err != nil {
			log.Fatalf("Failed to process the logs: %v", err)
		}

		fmt.Printf("Senders: %v\n", senders)

	default:
		log.Fatalf("Adapter %s is not supported", adapter)
	}
}
