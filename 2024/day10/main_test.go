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
	part1, part2 := solve(pkg.NewScanner(inputTest))
	assert.Equal(t, 36, part1)
	assert.Equal(t, 81, part2)
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(pkg.NewScanner(input))
	}
}
