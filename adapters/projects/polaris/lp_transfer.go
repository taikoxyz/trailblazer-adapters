package polaris

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/balancer_vault"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/erc20"
)

const (
	VaultAddress             = "0x723F40720836a03f1F20441FbDa3D109fc99640F" // Replace with the actual vault contract address
	logPoolBalanceChangedSig = "PoolBalanceChanged(bytes32,address,address[],int256[],uint256[])"
)

type LPTransferIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewLPTransferIndexer(client *ethclient.Client, addresses []common.Address) *LPTransferIndexer {
	return &LPTransferIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.LPTransfer] = &LPTransferIndexer{}

func (indexer *LPTransferIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *LPTransferIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.LPTransfer, error) {
	var transfers []adapters.LPTransfer

	for _, l := range logs {
		if !indexer.isPoolBalanceChanged(l) {
			continue
		}

		transfer, err := indexer.processPoolBalanceChangedLog(ctx, l)
		if err != nil {
			return nil, errors.Join(err, errors.New("processing PoolBalanceChanged log"))
		}

		if transfer != nil {
			transfers = append(transfers, *transfer)
		}
	}

	return transfers, nil
}

func (indexer *LPTransferIndexer) processPoolBalanceChangedLog(ctx context.Context, l types.Log) (*adapters.LPTransfer, error) {
	event, err := indexer.parsePoolBalanceChangedEvent(l)
	if err != nil {
		return nil, errors.Join(err, errors.New("parsing PoolBalanceChanged event"))
	}

	// Check if both deltas are positive
	if len(event.Deltas) != 2 {
		// Assuming we are only interested in events with exactly two tokens
		return nil, nil // Skip this event
	}

	if event.Deltas[0].Sign() <= 0 || event.Deltas[1].Sign() <= 0 {
		// One or both deltas are not positive, skip this event
		return nil, nil
	}

	block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
	if err != nil {
		return nil, errors.Join(err, errors.New("fetching block"))
	}

	// Fetch token decimals for both tokens
	token0Decimals, err := indexer.fetchTokenDecimals(ctx, event.Tokens[0], block.Number())
	if err != nil {
		return nil, errors.Join(err, errors.New("fetching token0 decimals"))
	}

	token1Decimals, err := indexer.fetchTokenDecimals(ctx, event.Tokens[1], block.Number())
	if err != nil {
		return nil, errors.Join(err, errors.New("fetching token1 decimals"))
	}

	// Create an LPTransfer record with the vault address as the recipient
	vaultAddress := common.HexToAddress(VaultAddress)

	transfer := adapters.LPTransfer{
		From:           event.LiquidityProvider,
		To:             vaultAddress,
		Token0Amount:   event.Deltas[0],
		Token0Decimals: token0Decimals,
		Token0:         event.Tokens[0],
		Token1Amount:   event.Deltas[1],
		Token1Decimals: token1Decimals,
		Token1:         event.Tokens[1],
		Time:           block.Time(),
		BlockNumber:    block.NumberU64(),
		TxHash:         l.TxHash,
	}

	return &transfer, nil
}

type PoolBalanceChangedEvent struct {
	PoolId             [32]byte
	LiquidityProvider  common.Address
	Tokens             []common.Address
	Deltas             []*big.Int
	ProtocolFeeAmounts []*big.Int
}

func (indexer *LPTransferIndexer) parsePoolBalanceChangedEvent(l types.Log) (*PoolBalanceChangedEvent, error) {
	if len(l.Topics) != 3 {
		return nil, errors.New("incorrect number of topics")
	}

	// The PoolBalanceChanged event is indexed on PoolId and LiquidityProvider
	var poolId [32]byte
	copy(poolId[:], l.Topics[1].Bytes())

	liquidityProvider := common.BytesToAddress(l.Topics[2].Bytes()[12:])

	vaultABI, err := abi.JSON(strings.NewReader(balancer_vault.BalancerVaultABI))
	if err != nil {
		return nil, errors.Join(err, errors.New("parsing ABI"))
	}

	var eventData struct {
		Tokens             []common.Address
		Deltas             []*big.Int
		ProtocolFeeAmounts []*big.Int
	}

	// Unpack the event data
	err = vaultABI.UnpackIntoInterface(&eventData, "PoolBalanceChanged", l.Data)
	if err != nil {
		return nil, errors.Join(err, errors.New("unpacking event data"))
	}

	event := &PoolBalanceChangedEvent{
		PoolId:             poolId,
		LiquidityProvider:  liquidityProvider,
		Tokens:             eventData.Tokens,
		Deltas:             eventData.Deltas,
		ProtocolFeeAmounts: eventData.ProtocolFeeAmounts,
	}

	return event, nil
}

func (indexer *LPTransferIndexer) fetchTokenDecimals(ctx context.Context, token common.Address, blockNumber *big.Int) (uint8, error) {
	erc20ABI, err := abi.JSON(strings.NewReader(erc20.Erc20ABI))
	if err != nil {
		return 0, errors.Join(err, errors.New("parsing ERC20 ABI"))
	}

	data, err := erc20ABI.Pack("decimals")
	if err != nil {
		return 0, errors.Join(err, errors.New("packing decimals call"))
	}

	callMsg := ethereum.CallMsg{
		To:   &token,
		Data: data,
	}

	result, err := indexer.client.CallContract(ctx, callMsg, blockNumber)
	if err != nil {
		return 0, errors.Join(err, errors.New("calling decimals method"))
	}

	var decimals uint8
	err = erc20ABI.UnpackIntoInterface(&decimals, "decimals", result)
	if err != nil {
		return 0, errors.Join(err, errors.New("unpacking decimals result"))
	}

	return decimals, nil
}

func (indexer *LPTransferIndexer) isPoolBalanceChanged(l types.Log) bool {
	eventSignatureHash := crypto.Keccak256Hash([]byte(logPoolBalanceChangedSig))
	return l.Topics[0].Hex() == eventSignatureHash.Hex()
}
