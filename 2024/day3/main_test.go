package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 179834255, solve(input, false))
	assert.Equal(t, 80570939, solve(input, true))
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
