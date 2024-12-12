package main

import (
	"bufio"
	"bytes"
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var inputTest []byte

func Test(t *testing.T) {
	reader := bytes.NewReader(inputTest)
	sum, sim := solve(bufio.NewScanner(reader))

	assert.Equal(t, 11, sum)
	assert.Equal(t, 31, sim)
}

func Benchmark(b *testing.B) {
	reader := bytes.NewReader(input)

	for i := 0; i < b.N; i++ {
		_, _ = reader.Seek(0, 0)
		solve(bufio.NewScanner(reader))
	}
}
