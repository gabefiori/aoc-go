package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	part1, part2 := solve(input)

	assert.Equal(t, 2573, part1)
	assert.Equal(t, 1850, part2)
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
