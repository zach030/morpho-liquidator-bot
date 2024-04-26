package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	VIRTUAL_SHARES                   = big.NewInt(1e6) // 10**6
	VIRTUAL_ASSETS                   = big.NewInt(1)
	WAD                              = pow10(18)
	ORACLE_PRICE_SCALE, _            = new(big.Int).SetString("1000000000000000000000000000000000000", 10)
	MAX_LIQUIDATION_INCENTIVE_FACTOR = new(big.Int).SetUint64(1.15e18)
	LIQUIDATION_CURSOR               = new(big.Int).SetUint64(0.3e18)
)

// EventTopic returns the keccak256 hash of `str`.
func EventTopic(str string) common.Hash {
	return crypto.Keccak256Hash([]byte(str))
}

// GetSeizedAmtByShares calculates the amount of assets seized by repaying `repaidShares` of a loan.
func GetSeizedAmtByShares(repaidShares *big.Int, market *MarketInfo) *big.Int {
	oneMinusLltv := new(big.Int).Sub(WAD, market.Params.Lltv)
	incentiveCalc := new(big.Int).Sub(WAD, wMulDown(LIQUIDATION_CURSOR, oneMinusLltv))
	liquidationIncentiveFactor := min(MAX_LIQUIDATION_INCENTIVE_FACTOR, wDivDown(WAD, incentiveCalc))
	a := toAssetsDown(repaidShares, market.MarketState.TotalBorrowAssets, market.MarketState.TotalBorrowShares)
	b := wMulDown(a, liquidationIncentiveFactor)
	return mulDivDown(b, market.CollateralPrice, ORACLE_PRICE_SCALE)
}

// pow10 returns 10 raised to the power of `exponent`.
func pow10(exponent int) *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(exponent)), nil)
}

// toAssetsUp calculates the value of `shares` quoted in assets, rounding up.
func toAssetsUp(shares, totalAssets, totalShares *big.Int) *big.Int {
	return mulDivUp(shares, new(big.Int).Add(totalAssets, VIRTUAL_ASSETS), new(big.Int).Add(totalShares, VIRTUAL_SHARES))
}

// mulDivUp multiplies `x` by `y` and divides by `d`, rounding up.
func mulDivUp(x, y, d *big.Int) *big.Int {
	mul := new(big.Int).Mul(x, y)
	add := new(big.Int).Add(mul, new(big.Int).Sub(d, big.NewInt(1)))
	return new(big.Int).Div(add, d)
}

// wMulDown multiplies `x` by `y` and divides by `WAD` rounding down.
func wMulDown(x, y *big.Int) *big.Int {
	return mulDivDown(x, y, WAD)
}

// wDivDown divides `x` by `y` and multiplies by `WAD` rounding down.
func wDivDown(x, y *big.Int) *big.Int {
	return mulDivDown(x, WAD, y)
}

// wDivUp divides `x` by `y` and multiplies by `WAD` rounding up.
func wDivUp(x, y *big.Int) *big.Int {
	return mulDivUp(x, WAD, y)
}

// mulDivDown multiplies `x` by `y` and divides by `d`, rounding down.
func mulDivDown(x, y, d *big.Int) *big.Int {
	return new(big.Int).Div(new(big.Int).Mul(x, y), d)
}

// min returns the minimum of `a` and `b`.
func min(a, b *big.Int) *big.Int {
	if a.Cmp(b) < 0 {
		return a
	}
	return b
}

// toAssetsDown calculates the value of `shares` quoted in assets, rounding down.
func toAssetsDown(shares, totalAssets, totalShares *big.Int) *big.Int {
	return mulDivDown(shares, new(big.Int).Add(totalAssets, VIRTUAL_ASSETS), new(big.Int).Add(totalShares, VIRTUAL_SHARES))
}
