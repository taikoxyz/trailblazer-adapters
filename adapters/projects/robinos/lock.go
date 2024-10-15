package symmetric

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/balancer_token"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/balancer_vault"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/symmetric"
)

const (
	RobinosAddress                = "0x9EF9c57754eD079D750016b802DcCD45d0AB66f8"
	LogRewardDistributedSignature = "RewardDistributed(string,uint256[],address[])"
)

var SelectedMultiplierEvents = []string{"Prediction - $TAIKO price on Oct 20", "Serie A 24/25 - Juventus v. Lazio", "EPL 24/25 - Liverpool v. Chelsea", "EPL 24/25 - Man Utd v. Brentford"}

type RobinosIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewMultiplierEventIndexer(client *ethclient.Client, addresses []common.Address) *RobinosIndexer {
	return &RobinosIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.RewardIndexer[adapters.Robinos] = &RobinosIndexer{}

func (indexer *RobinosIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *LockIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Lock, error) {
	var locks []adapters.Lock

	for _, l := range logs {
		if !indexer.isDeposit(l) {
			continue
		}

		lock, err := indexer.processDepositLog(ctx, l)
		if err != nil {
			return nil, errors.Join(err, errors.New("processing deposit log"))
		}

		locks = append(locks, lock...)
	}

	return locks, nil
}

func (indexer *LockIndexer) processDepositLog(ctx context.Context, l types.Log) ([]adapters.Lock, error) {
	depositEvent, err := indexer.parseDepositEvent(l)
	if err != nil {
		return nil, errors.Join(err, errors.New("parsing deposit event"))
	}

	user := common.BytesToAddress(l.Topics[1].Bytes()[12:])
	lockEndTime := l.Topics[2].Big()
	duration := new(big.Int).Sub(lockEndTime, depositEvent.Ts)

	block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
	if err != nil {
		return nil, errors.Join(err, errors.New("fetching block"))
	}

	symmetricContract, err := symmetric.NewSymmetric(l.Address, indexer.client)
	if err != nil {
		return nil, errors.Join(err, errors.New("creating symmetric contract"))
	}

	token, err := symmetricContract.Token(&bind.CallOpts{BlockNumber: block.Number()})
	if err != nil {
		return nil, errors.Join(err, errors.New("fetching token"))
	}

	balancerToken, err := balancer_token.NewBalancerToken(token, indexer.client)
	if err != nil {
		return nil, errors.Join(err, errors.New("creating balancer token"))
	}

	totalSupply, poolId, vault, err := indexer.fetchBalancerTokenInfo(balancerToken, block.Number())
	if err != nil {
		return nil, errors.Join(err, errors.New("fetching balancer token info"))
	}

	balancerVault, err := balancer_vault.NewBalancerVault(vault, indexer.client)
	if err != nil {
		return nil, errors.Join(err, errors.New("creating balancer vault"))
	}

	poolTokens, err := balancerVault.GetPoolTokens(&bind.CallOpts{BlockNumber: block.Number()}, poolId)
	if err != nil {
		return nil, errors.Join(err, errors.New("fetching pool tokens"))
	}

	return indexer.createLocks(user, duration, depositEvent.Value, totalSupply, poolTokens, block, l.TxHash), nil
}

func (indexer *LockIndexer) parseDepositEvent(l types.Log) (*struct {
	Value *big.Int
	Type  *big.Int
	Ts    *big.Int
}, error) {
	var depositEvent struct {
		Value *big.Int
		Type  *big.Int
		Ts    *big.Int
	}

	symmetricABI, err := abi.JSON(strings.NewReader(symmetric.SymmetricABI))
	if err != nil {
		return nil, errors.Join(err, errors.New("parsing ABI"))
	}

	err = symmetricABI.UnpackIntoInterface(&depositEvent, "Deposit", l.Data)
	if err != nil {
		return nil, errors.Join(err, errors.New("unpacking event data"))
	}

	return &depositEvent, nil
}

func (indexer *LockIndexer) fetchBalancerTokenInfo(balancerToken *balancer_token.BalancerToken, blockNumber *big.Int) (*big.Int, [32]byte, common.Address, error) {
	opts := &bind.CallOpts{BlockNumber: blockNumber}

	totalSupply, err := balancerToken.TotalSupply(opts)
	if err != nil {
		return nil, [32]byte{}, common.Address{}, errors.Join(err, errors.New("fetching total supply"))
	}

	poolId, err := balancerToken.GetPoolId(opts)
	if err != nil {
		return nil, [32]byte{}, common.Address{}, errors.Join(err, errors.New("fetching pool ID"))
	}

	vault, err := balancerToken.GetVault(opts)
	if err != nil {
		return nil, [32]byte{}, common.Address{}, errors.Join(err, errors.New("fetching vault"))
	}

	return totalSupply, poolId, vault, nil
}

func (indexer *LockIndexer) createLocks(user common.Address, duration, depositValue *big.Int, totalSupply *big.Int, poolTokens struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}, block *types.Block, txHash common.Hash) []adapters.Lock {
	var locks []adapters.Lock

	for i, token := range poolTokens.Tokens {
		tokenAmount := poolTokens.Balances[i]
		userAmount := new(big.Int).Mul(tokenAmount, depositValue)
		userAmount.Div(userAmount, totalSupply)

		lock := adapters.Lock{
			User:          user,
			TokenAmount:   userAmount,
			TokenDecimals: adapters.TaikoTokenDecimals,
			Token:         token,
			Duration:      duration.Uint64(),
			BlockTime:     block.Time(),
			BlockNumber:   block.NumberU64(),
			TxHash:        txHash,
		}
		locks = append(locks, lock)
	}

	return locks
}

func (indexer *LockIndexer) isDeposit(l types.Log) bool {
	return l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logDepositSignature)).Hex()
}
