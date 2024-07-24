package cmd

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
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
	var adapterOptions = []string{"NewTransactionSender", "DotTaikoIndexer", "OrderFulfilledIndexer", "NewSaleIndexer"}
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
		processNewTransactionSender(client, blockNumber)
	case "OrderFulfilledIndexer":
		processOrderFulfilledIndexer(client, blockNumber)
	case "DotTaikoIndexer":
		processDotTaikoIndexer(client, blockNumber)
	case "NewSaleIndexer":
		processNewSaleIndexer(client, blockNumber)
	default:
		log.Fatalf("Adapter %s is not supported", adapter)
	}
}
