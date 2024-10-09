package optional

import "math/big"

func FromBigInt(a *big.Int) big.Int {
	if a == nil {
		return *big.NewInt(0)
	}

	return *a
}
