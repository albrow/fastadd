package fastadd

import (
	"errors"
	"math"
)

// OverflowError is returned whenever a number overflows int64.
var OverflowError = errors.New("integer overflow")

// FastAdd is a "highly optimized" addition function which uses bitwise
// operations.
//
// It returns an OverflowError if the result of x + y overflows int64. (i.e. if
// it is less than math.MinInt64 or greater than math.MaxInt64).
func FastAdd(x int, y int) (int, error) {
	halfAdd := x ^ y
	carry := (x & y) << 1
	result := halfAdd + carry

	if result > math.MaxInt64 || result < math.MinInt64 {
		return 0, OverflowError
	}

	return result, nil
}
