package pkg

import (
	"bufio"
	"bytes"
)

func NewScanner(input []byte) *bufio.Scanner {
	r := bytes.NewReader(input)
	return bufio.NewScanner(r)
}
