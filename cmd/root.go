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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "trailblazer-adapters",
	Short: "Trailblazer Adapters CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := promptUser(); err != nil {
			log.Fatalf("Prompt failed: %v", err)
		}
		if err := executeCommand(); err != nil {
			log.Fatalf("Execution failed: %v", err)
		}
	},
}

func init() {
	cobra.OnInitialize()
}

func promptUser() error {
	var adapterOptions = []string{
		"Izumi", "RitsuLP", "NewTransactionSender", "NftDeployed", "GamingWhitelist", "DotTaikoIndexer",
		"OrderFulfilledIndexer", "NewSaleIndexer", "ContractDeployed",
		"CollectionCreated", "TokenSold",
	}

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

	if err := survey.Ask(qs, &answers); err != nil {
		return err
	}

	adapter = answers.Adapter
	blockNumber = answers.BlockNumber
	return nil
}

func executeCommand() error {
	client, err := ethclient.Dial("https://rpc.taiko.xyz")
	if err != nil {
		return fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	switch adapter {
	case "RitsuLP":
		return processRitsuLPIndexer(client, blockNumber)
	case "Izumi":
		return processIziLPIndexer(client, blockNumber)
	case "NewTransactionSender":
		return processNewTransactionSender(client, blockNumber)
	case "NftDeployed":
		return processNewNftDeployed(client, blockNumber)
	case "GamingWhitelist":
		return processNewGamingWhitelist(client, blockNumber)
	case "OrderFulfilledIndexer":
		return processOrderFulfilledIndexer(client, blockNumber)
	case "DotTaikoIndexer":
		return processDotTaikoIndexer(client, blockNumber)
	case "NewSaleIndexer":
		return processNewSaleIndexer(client, blockNumber)
	case "ContractDeployed":
		return processContractDeployedIndexer(client, blockNumber)
	case "CollectionCreated":
		return processCollectionCreatedIndexer(client, blockNumber)
	case "TokenSold":
		return processTokenSoldIndexer(client, blockNumber)
	default:
		return fmt.Errorf("adapter %s is not supported", adapter)
	}
}
