package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	part1, part2 := solve(bytes.NewBuffer(input))

	assert.Equal(t, 4637, part1)
	assert.Equal(t, 6370, part2)
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(bytes.NewBuffer(input))
	}
}
