package pkg

import (
	"golang.org/x/exp/constraints"
	"strconv"
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

// ParseInt ignores conversion errors
func ParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
