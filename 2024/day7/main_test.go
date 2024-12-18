package main

import (
	"testing"

	_ "embed"

	"github.com/gabefiori/aoc-go/pkg"
	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var inputTest []byte

func Test(t *testing.T) {
	assert.Equal(t, 3749, solve(pkg.NewScanner(inputTest), []int{add, mult}))
	assert.Equal(t, 11387, solve(pkg.NewScanner(inputTest), []int{add, mult, concat}))
}

func Benchmark(b *testing.B) {
	b.Run("Part 1", func(b *testing.B) {
		ops := []int{add, mult}
		for i := 0; i < b.N; i++ {
			solve(pkg.NewScanner(input), ops)
		}
	})

	b.Run("Part 2", func(b *testing.B) {
		ops := []int{add, mult, concat}
		for i := 0; i < b.N; i++ {
			solve(pkg.NewScanner(input), ops)
		}
	})
}
