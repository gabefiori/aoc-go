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
	assert.Equal(t, 3749, solve(pkg.NewScanner(inputTest), []int{add, mult}))
	assert.Equal(t, 11387, solve(pkg.NewScanner(inputTest), []int{add, mult, concat}))
}

func Benchmark(b *testing.B) {
	b.Run("Part 1", func(b *testing.B) {
		reader := bytes.NewReader(input)
		for i := 0; i < b.N; i++ {
			_, _ = reader.Seek(0, 0)
			solve(bufio.NewScanner(reader), []int{add, mult})
		}
	})

	b.Run("Part 2", func(b *testing.B) {
		reader := bytes.NewReader(input)
		for i := 0; i < b.N; i++ {
			_, _ = reader.Seek(0, 0)
			solve(bufio.NewScanner(reader), []int{add, mult, concat})
		}
	})
}
