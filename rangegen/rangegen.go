// Package rangegen provides common helper functions for working with ranges of values.
package rangegen

import (
	"math/big"
)

var one = big.NewInt(1)

func maxVal(numBytes uint) *big.Int {
	limit := new(big.Int).Lsh(one, 8*numBytes)
	return limit.Sub(limit, one)
}
