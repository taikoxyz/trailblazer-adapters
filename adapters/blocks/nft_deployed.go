package blocks

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/erc165"
)

type NftDeployedIndexer struct{}

// NewNftDeployedIndexer creates a new NftDeployedIndexer.
func NewNftDeployedIndexer() *NftDeployedIndexer {
	return &NftDeployedIndexer{}
}

const (
	ERC721InterfaceID  = "0x80ac58cd" // ERC721 interface ID
	ERC1155InterfaceID = "0xd9b67a26" // ERC1155 interface ID
)

func supportsInterface(contractAddress common.Address, client *ethclient.Client, interfaceID string) (bool, error) {
	parsedABI, err := abi.JSON(strings.NewReader(erc165.ABI))
	if err != nil {
		return false, err
	}

	interfaceIDBytes := common.FromHex(interfaceID)
	if len(interfaceIDBytes) != 4 {
		return false, fmt.Errorf("invalid interface ID length")
	}

	data, err := parsedABI.Pack("supportsInterface", [4]byte{interfaceIDBytes[0], interfaceIDBytes[1], interfaceIDBytes[2], interfaceIDBytes[3]})
	if err != nil {
		return false, err
	}

	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return false, err
	}

	var supports bool
	err = parsedABI.UnpackIntoInterface(&supports, "supportsInterface", result)
	if err != nil {
		return false, err
	}

	return supports, nil
}

func (indexer *NftDeployedIndexer) ProcessBlock(ctx context.Context, block *types.Block, client *ethclient.Client) ([]common.Address, error) {
	txs := block.Transactions()
	var result []common.Address

	for _, tx := range txs {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return nil, err
		}

		sender, err := client.TransactionSender(ctx, tx, block.Hash(), receipt.TransactionIndex)
		if err != nil {
			return nil, err
		}

		to := tx.To()
		if to != nil {
			continue
		}

		isErc721, err := supportsInterface(receipt.ContractAddress, client, ERC721InterfaceID)
		if err != nil {
			return nil, err
		}

		if isErc721 {
			result = append(result, sender)
		}

		isErc1155, err := supportsInterface(receipt.ContractAddress, client, ERC1155InterfaceID)
		if err != nil {
			return nil, err
		}

		if isErc1155 {
			result = append(result, sender)
		}
	}

	return result, nil
}
