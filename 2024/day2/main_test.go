package main

import (
	"bufio"
	"bytes"
	"testing"

	_ "embed"

	"github.com/gabefiori/aoc-go/pkg"
	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var inputTest []byte

func Test(t *testing.T) {
	assert.Equal(t, 2, solve(pkg.NewScanner(inputTest), false))
	assert.Equal(t, 4, solve(pkg.NewScanner(inputTest), true))
}

func Benchmark(b *testing.B) {
	b.Run("Part 1", func(b *testing.B) {
		reader := bytes.NewReader(input)
		for i := 0; i < b.N; i++ {
			_, _ = reader.Seek(0, 0)
			solve(bufio.NewScanner(reader), false)
		}
	})

	b.Run("Part 2", func(b *testing.B) {
		reader := bytes.NewReader(input)
		for i := 0; i < b.N; i++ {
			_, _ = reader.Seek(0, 0)
			solve(bufio.NewScanner(reader), true)
		}
	})
}
