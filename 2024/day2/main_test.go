package main

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/gabefiori/aoc-go/pkg"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 402, solve(pkg.NewScanner(input), 0))
	assert.Equal(t, 455, solve(pkg.NewScanner(input), 1))
}

func Benchmark(b *testing.B) {
	b.Run("Part 1", func(b *testing.B) {
		reader := bytes.NewReader(input)
		for i := 0; i < b.N; i++ {
			_, _ = reader.Seek(0, 0)
			solve(bufio.NewScanner(reader), 0)
		}
	})

	b.Run("Part 2", func(b *testing.B) {
		reader := bytes.NewReader(input)
		for i := 0; i < b.N; i++ {
			_, _ = reader.Seek(0, 0)
			solve(bufio.NewScanner(reader), 1)
		}
	})
}
