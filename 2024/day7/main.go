package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/gabefiori/aoc-go/pkg"
)

const (
	add int = iota
	mult
	concat
)

//go:embed input.txt
var input []byte

func solve(sc *bufio.Scanner, operators []int) int {
	var sum, target int
	var part, comb, vals []int

	for sc.Scan() {
		line := sc.Text()
		sep := strings.Index(line, ":")
		target = pkg.ParseInt(line[:sep])
		vals = pkg.ParseInts(strings.Split(line[sep+2:], " "), vals)

		part, comb = part[:0], comb[:0]
		genOperators(len(vals)-1, part, operators, &comb)

		for ops := range slices.Chunk(comb, len(vals)-1) {
			t := vals[0]
			for i, val := range vals[1:] {
				switch ops[i] {
				case add:
					t += val
				case mult:
					t *= val
				case concat:
					t = pkg.ParseInt(strconv.Itoa(t) + strconv.Itoa(val))
				}
			}

			if t == target {
				sum += t
				break
			}
		}
	}

	return sum
}

func genOperators(n int, curr, ops []int, res *[]int) {
	if len(curr) == n {
		*res = append(*res, curr...)
		return
	}

	for _, op := range ops {
		genOperators(n, append(curr, op), ops, res)
	}
}

func main() {
	part1 := solve(pkg.NewScanner(input), []int{add, mult})
	part2 := solve(pkg.NewScanner(input), []int{add, mult, concat})
	fmt.Println(part1, part2)
}
