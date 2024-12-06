package pkg

import (
	"cmp"
	"strconv"

	"golang.org/x/exp/constraints"
)

func AbsDiff[T constraints.Integer](a, b T) T {
	if a > b {
		return a - b
	}

	return b - a
}

func Abs[T constraints.Signed](n T) T {
	if n < 0 {
		return -n
	}

	return n
}

func Clamp[T cmp.Ordered](v, min, max T) T {
	if v < min {
		return min
	}

	if v > max {
		return max
	}

	return v
}

// ParseInt ignores conversion errors
func ParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
