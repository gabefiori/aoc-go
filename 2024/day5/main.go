package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"

	"github.com/gabefiori/aoc-go/pkg"
)

//go:embed input.txt
var input []byte

func solve(buf *bytes.Buffer) (int, int) {
	var sum, sumFix int

	rules := make(map[int][]int, 50)

	for {
		line, err := buf.ReadString('\n')
		if err != nil || line == "\n" {
			break
		}

		page, rule := pkg.ParseInt(line[:2]), pkg.ParseInt(line[3:5])
		rules[page] = append(rules[page], rule)
	}

	updates := make([]int, 0, 100)

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			break
		}

		var mid, count int

		parts := strings.Split(line[:len(line)-1], ",")
		updates = pkg.ParseInts(parts, updates)

		for i := 1; i < len(updates); i++ {
			rule, exists := rules[updates[i]]
			if !exists {
				continue
			}

			found := contains(rule, updates[:i])
			if found >= 0 {
				updates[i], updates[found] = updates[found], updates[i]
				i = found - 1

				count++
			}

			mid = updates[len(updates)/2]
		}

		if count == 0 {
			sum += mid
		} else {
			sumFix += mid
		}
	}

	return sum, sumFix
}

func contains(s, v []int) int {
	for _, x := range s {
		for i, y := range v {
			if x == y {
				return i
			}
		}
	}
	return -1
}

func main() {
	part1, part2 := solve(bytes.NewBuffer(input))
	fmt.Println(part1, part2)
}
