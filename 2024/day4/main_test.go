package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var inputTest string

func Test(t *testing.T) {
	part1, part2 := solve(inputTest)

	assert.Equal(t, 18, part1)
	assert.Equal(t, 9, part2)
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
