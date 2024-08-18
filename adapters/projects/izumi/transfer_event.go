package izumi

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/erc20"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/izipool"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/izumi"
)

var (
	logTransferSigHash = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
)

// TransferIndexer is an implementation of LogsIndexer for ERC20 transfer logs.
type TransferIndexer struct {
	token     common.Address
	whitelist map[string]struct{}
}

// NewTransferIndexer creates a new TransferIndexer.
func NewTransferIndexer(token common.Address, whitelist map[string]struct{}) *TransferIndexer {
	return &TransferIndexer{
		token:     token,
		whitelist: whitelist,
	}
}

func (indexer *TransferIndexer) Address() common.Address {
	return indexer.token
}

// IndexLogs processes logs for ERC20 transfers.
func (indexer *TransferIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.LPTransfer, error) {
	var result []adapters.LPTransfer
	for _, vLog := range logs {
		if !isERC721Transfer(vLog) {
			continue
		}
		transferData, err := indexer.ProcessLog(ctx, chainID, client, vLog)
		if err != nil {
			return nil, err
		}
		if transferData != nil {
			result = append(result, *transferData)
		}
	}
	return result, nil
}

func isERC721Transfer(vLog types.Log) bool {
	return len(vLog.Topics) == 4 && vLog.Topics[0].Hex() == logTransferSigHash.Hex()
}

func (indexer *TransferIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.LPTransfer, error) {
	// Extract "from" and "to" addresses from the log
	to := common.BytesToAddress(vLog.Topics[2].Bytes()[12:])
	_, exists := indexer.whitelist[to.Hex()]
	if !exists {
		return nil, nil
	}
	from := common.BytesToAddress(vLog.Topics[1].Bytes()[12:])
	tokenID := vLog.Topics[3].Big()

	// Fetch the block details
	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	// Initialize the LiquidityManager contract caller
	liquidityManager, err := izumi.NewIzumiCaller(indexer.token, client)
	if err != nil {
		return nil, err
	}
	// Fetch liquidity details using the Token ID (NFT)
	liquidity, err := liquidityManager.Liquidities(&bind.CallOpts{
		BlockNumber: block.Number(),
	}, tokenID)
	if err != nil {
		return nil, err
	}

	// Fetch pool metadata using the pool ID
	poolMeta, err := liquidityManager.PoolMetas(&bind.CallOpts{
		BlockNumber: block.Number(),
	}, liquidity.PoolId)
	if err != nil {
		return nil, err
	}

	// Fetch token addresses and decimals for the pool
	token0Address := poolMeta.TokenX
	token1Address := poolMeta.TokenY

	_, token0Decimals, err := fetchTokenDetails(token0Address, client)
	if err != nil {
		return nil, err
	}

	_, token1Decimals, err := fetchTokenDetails(token1Address, client)
	if err != nil {
		return nil, err
	}

	// Calculate the amount of tokens in the liquidity position
	token0Amount, token1Amount, err := calculateLiquidityAmounts(liquidityManager, liquidity, poolMeta, block, client)
	if err != nil {
		return nil, err
	}

	// Return the LPTransfer struct with calculated values
	return &adapters.LPTransfer{
		From:           from,
		To:             to,
		Token0Amount:   token0Amount,
		Token0Decimals: token0Decimals,
		Token0:         token0Address,
		Token1Amount:   token1Amount,
		Token1Decimals: token1Decimals,
		Token1:         token1Address,
		Time:           block.Time(),
		BlockNumber:    block.Number().Uint64(),
	}, nil
}

// Helper function to fetch token details (caller and decimals)
func fetchTokenDetails(tokenAddress common.Address, client *ethclient.Client) (token *erc20.Erc20Caller, decimals uint8, err error) {
	token, err = erc20.NewErc20Caller(tokenAddress, client)
	if err != nil {
		return
	}

	decimals, err = token.Decimals(nil)
	return
}

// Helper function to calculate the amount of tokens in the liquidity position
func calculateLiquidityAmounts(caller *izumi.IzumiCaller, liquidity struct {
	LeftPt           *big.Int
	RightPt          *big.Int
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	RemainTokenX     *big.Int
	RemainTokenY     *big.Int
	PoolId           *big.Int
}, poolMeta struct {
	TokenX common.Address
	TokenY common.Address
	Fee    *big.Int
}, block *types.Block, client *ethclient.Client) (token0Amount, token1Amount *big.Int, err error) {
	poolAddress, err := caller.Pool(&bind.CallOpts{
		BlockNumber: block.Number(),
	}, poolMeta.TokenX, poolMeta.TokenY, poolMeta.Fee)

	if poolAddress == adapters.ZeroAddress || err != nil {
		return nil, nil, fmt.Errorf("pool not found")
	}

	pool, err := izipool.NewIziPoolCaller(poolAddress, client)
	if err != nil {
		return nil, nil, err
	}
	// Get current pool price and point
	state, err := getPoolPrice(pool, block)
	if err != nil {
		return nil, nil, err
	}

	amountX, amountY := getLiquidityValue(liquidity, *state)

	return amountX, amountY, nil
}

// Helper function to get the current price and point for the pool
func getPoolPrice(pool *izipool.IziPoolCaller, block *types.Block) (*struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}, error) {
	state, err := pool.State(&bind.CallOpts{
		BlockNumber: block.Number(),
	})
	if err != nil {
		return nil, err
	}

	return &state, nil
}
func pow(base *big.Float, exp *big.Int) *big.Float {
	result := new(big.Float).SetPrec(200).SetFloat64(1.0) // Initialize result as 1.0
	baseCopy := new(big.Float).SetPrec(200).Copy(base)    // Copy base to avoid modifying original

	expCopy := new(big.Int).Set(exp) // Copy exp to avoid modifying original
	zero := big.NewInt(0)

	for expCopy.Cmp(zero) > 0 { // While exp > 0
		result.Mul(result, baseCopy)
		expCopy.Sub(expCopy, big.NewInt(1))
	}

	return result
}

// Function to calculate sqrt(1.0001^point) using big.Float
func point2PoolPriceUndecimalSqrt(point *big.Int) *big.Float {
	constant := new(big.Float).SetPrec(200).SetFloat64(1.0001)

	// Determine if the point is negative
	isNegative := point.Sign() < 0

	// Take the absolute value of the point
	absPoint := new(big.Int).Abs(point)

	// Compute 1.0001^abs(point)
	expResult := pow(constant, absPoint)

	// Take the square root of the result
	sqrtResult := new(big.Float).SetPrec(200).Sqrt(expResult)

	// If the original point was negative, take the reciprocal of the result
	if isNegative {
		one := new(big.Float).SetPrec(200).SetFloat64(1.0)
		sqrtResult.Quo(one, sqrtResult)
	}

	return sqrtResult
}

// Function to calculate amountY using big.Float
func getAmountY(liquidity *big.Int, sqrtPriceL, sqrtPriceR, sqrtRate *big.Float, upper bool) *big.Int {
	numerator := new(big.Float).Sub(sqrtPriceR, sqrtPriceL)
	denominator := new(big.Float).Sub(sqrtRate, big.NewFloat(1.0))

	amount := new(big.Float).SetPrec(200).SetInt(liquidity)
	amount.Mul(amount, numerator)
	amount.Quo(amount, denominator)

	result := new(big.Int)
	if upper {
		amount.Add(amount, big.NewFloat(0.5)) // Round up
	}
	amount.Int(result)

	return result
}

// Function to calculate amountY at a specific point using big.Float
func liquidity2AmountYAtPoint(liquidity *big.Int, sqrtPrice *big.Float, upper bool) *big.Int {
	amountY := new(big.Float).SetPrec(200).SetInt(liquidity)
	amountY.Mul(amountY, sqrtPrice)

	result := new(big.Int)
	if upper {
		amountY.Add(amountY, big.NewFloat(0.5)) // Round up
	}
	amountY.Int(result)

	return result
}

// Function to calculate amountX using big.Float
func getAmountX(liquidity *big.Int, leftPt, rightPt *big.Int, sqrtPriceR, sqrtRate *big.Float, upper bool) *big.Int {
	// Calculate sqrtPricePrPc = sqrtRate^(rightPt-leftPt+1)
	diff := new(big.Int).Sub(rightPt, leftPt)
	diff.Add(diff, big.NewInt(1))
	sqrtPricePrPc := pow(sqrtRate, diff)

	// Calculate sqrtPricePrPd = sqrtRate^(rightPt+1)
	rightPtPlusOne := new(big.Int).Add(rightPt, big.NewInt(1))
	sqrtPricePrPd := pow(sqrtRate, rightPtPlusOne)

	numerator := new(big.Float).Sub(sqrtPricePrPc, sqrtRate)
	denominator := new(big.Float).Sub(sqrtPricePrPd, sqrtPriceR)

	amount := new(big.Float).SetPrec(200).SetInt(liquidity)
	amount.Mul(amount, numerator)
	amount.Quo(amount, denominator)

	result := new(big.Int)
	if upper {
		amount.Add(amount, big.NewFloat(0.5)) // Round up
	}
	amount.Int(result)

	return result
}

// Function to calculate amountX at a specific point using big.Float
func liquidity2AmountXAtPoint(liquidity *big.Int, sqrtPrice *big.Float, upper bool) *big.Int {
	amountX := new(big.Float).SetPrec(200).SetInt(liquidity)
	amountX.Quo(amountX, sqrtPrice)

	result := new(big.Int)
	if upper {
		amountX.Add(amountX, big.NewFloat(0.5)) // Round up
	}
	amountX.Int(result)

	return result
}

// Helper function to calculate the min of two *big.Int values
func minBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) < 0 {
		return a
	}
	return b
}

// Helper function to calculate the max of two *big.Int values
func maxBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}

// Main function to calculate the liquidity value
func getLiquidityValue(liquidity struct {
	LeftPt           *big.Int
	RightPt          *big.Int
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	RemainTokenX     *big.Int
	RemainTokenY     *big.Int
	PoolId           *big.Int
}, state struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}) (amountX, amountY *big.Int) {
	amountX = big.NewInt(0)
	amountY = big.NewInt(0)
	sqrtRate := new(big.Float).SetPrec(200).SetFloat64(math.Sqrt(1.0001))

	// Compute amountY without currentPoint
	if liquidity.LeftPt.Cmp(state.CurrentPoint) < 0 {
		rightPt := minBigInt(state.CurrentPoint, liquidity.RightPt)
		sqrtPriceR := point2PoolPriceUndecimalSqrt(rightPt)
		sqrtPriceL := point2PoolPriceUndecimalSqrt(liquidity.LeftPt)
		amountY = getAmountY(liquidity.Liquidity, sqrtPriceL, sqrtPriceR, sqrtRate, false)
	}

	// Compute amountX without currentPoint
	if liquidity.RightPt.Cmp(new(big.Int).Add(state.CurrentPoint, big.NewInt(1))) > 0 {
		leftPt := maxBigInt(new(big.Int).Add(state.CurrentPoint, big.NewInt(1)), liquidity.LeftPt)
		sqrtPriceR := point2PoolPriceUndecimalSqrt(liquidity.RightPt)
		amountX = getAmountX(liquidity.Liquidity, leftPt, liquidity.RightPt, sqrtPriceR, sqrtRate, false)
	}

	// Compute amountX and amountY on currentPoint
	if liquidity.LeftPt.Cmp(state.CurrentPoint) <= 0 && liquidity.RightPt.Cmp(state.CurrentPoint) > 0 {
		liquidityValue := new(big.Int).Set(liquidity.Liquidity)
		maxLiquidityYAtCurrentPt := new(big.Int).Sub(state.Liquidity, state.LiquidityX)
		liquidityYAtCurrentPt := minBigInt(liquidityValue, maxLiquidityYAtCurrentPt)
		liquidityXAtCurrentPt := new(big.Int).Sub(liquidityValue, liquidityYAtCurrentPt)
		currentSqrtPrice := point2PoolPriceUndecimalSqrt(state.CurrentPoint)
		amountX.Add(amountX, liquidity2AmountXAtPoint(liquidityXAtCurrentPt, currentSqrtPrice, false))
		amountY.Add(amountY, liquidity2AmountYAtPoint(liquidityYAtCurrentPt, currentSqrtPrice, false))
	}

	return amountX, amountY
}
