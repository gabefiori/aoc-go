package pkg

import (
	"bufio"
	"bytes"
)

func NewScanner(input []byte) *bufio.Scanner {
	reader := bytes.NewReader(input)
	return bufio.NewScanner(reader)
}
