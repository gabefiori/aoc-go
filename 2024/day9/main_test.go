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
	assert.Equal(t, 1928, solve(bytes.NewReader(inputTest), blocks))
	assert.Equal(t, 2858, solve(bytes.NewReader(inputTest), files))
}

func Benchmark(b *testing.B) {
	b.Run("Part 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			solve(bytes.NewReader(input), blocks)
		}
	})

	b.Run("Part 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			solve(bytes.NewReader(input), files)
		}
	})
}
