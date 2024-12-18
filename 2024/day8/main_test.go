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
	assert.Equal(t, 14, solve(pkg.NewScanner(inputTest), false))
	assert.Equal(t, 34, solve(pkg.NewScanner(inputTest), true))
}

func Benchmark(b *testing.B) {
	b.Run("Part 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			solve(pkg.NewScanner(input), false)
		}
	})

	b.Run("Part 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			solve(pkg.NewScanner(input), true)
		}
	})
}
