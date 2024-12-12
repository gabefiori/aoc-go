package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var inputTest string

//go:embed input_test-2.txt
var inputTest2 string

func Test(t *testing.T) {
	assert.Equal(t, 161, solve(inputTest, false))
	assert.Equal(t, 48, solve(inputTest2, true))
}

func Benchmark(b *testing.B) {
	b.Run("Part 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			solve(input, false)
		}
	})

	b.Run("Part 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			solve(input, false)
		}
	})
}
