package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"slices"
)

//go:embed input.txt
var input []byte

func solve(r io.Reader, compactFn func([]int)) int {
	disk, buf := make([]int, 0), make([]byte, 2)
	curr := 0

	for {
		_, err := r.Read(buf)
		if err != nil {
			break
		}

		for i := 0; i < int(buf[0]-48); i++ {
			disk = append(disk, curr)
		}

		if buf[1] > 10 {
			for i := 0; i < int(buf[1]-48); i++ {
				disk = append(disk, -1)
			}
		}

		curr++
	}

	compactFn(disk)

	checksum := 0
	for i, v := range disk {
		if v == -1 {
			v = 0
		}

		checksum += i * v
	}

	return checksum
}

func blocks(d []int) {
	i, j := 0, len(d)-1
	for i <= j {
		if d[i] > -1 {
			i++
		} else {
			d[i], d[j] = d[j], 0
			j--
		}
	}
}

func files(d []int) {
	slots := make([]int, 0, 9)
	i := len(d) - 1

	for i >= 0 {
		if d[i] == -1 {
			i--
			continue
		}

		block := slices.Index(d, d[i])
		slot := slices.Index(d, -1)

		for j := slot; j < len(d) && j < block; j++ {
			if d[j] != -1 {
				continue
			}

			slots = slots[:0]
			for j < block && d[j] == -1 {
				slots = append(slots, j)
				j++
			}

			if len(slots) >= i+1-block {
				for k := 0; k <= i-block; k++ {
					d[slots[k]], d[block+k] = d[block+k], 0
				}

				break
			}
		}

		i = block - 1
	}
}

func main() {
	part1 := solve(bytes.NewReader(input), blocks)
	part2 := solve(bytes.NewReader(input), files)
	fmt.Println(part1, part2)
}
