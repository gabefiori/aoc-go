package main

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	reader := bytes.NewReader(input)
	sum, sim := solve(bufio.NewScanner(reader))

	assert.Equal(t, 2769675, sum)
	assert.Equal(t, 24643097, sim)
}

func Benchmark(b *testing.B) {
	reader := bytes.NewReader(input)

	for i := 0; i < b.N; i++ {
		_, _ = reader.Seek(0, 0)
		solve(bufio.NewScanner(reader))
	}
}
