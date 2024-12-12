package main

import (
	"bytes"
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var inputTest []byte

func Test(t *testing.T) {
	part1, part2 := solve(bytes.NewBuffer(inputTest))

	assert.Equal(t, 143, part1)
	assert.Equal(t, 123, part2)
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(bytes.NewBuffer(input))
	}
}
