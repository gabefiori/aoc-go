package main

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"

	"github.com/gabefiori/aoc-go/pkg"
)

const sep = " "

const (
	ordAsc int = iota
	ordDesc
)

//go:embed input.txt
var input []byte

func solve(sc *bufio.Scanner, tolerance int) int {
	var safe int
	lvl := make([]int, 0, 10)

	for sc.Scan() {
		lvl = pkg.ParseInts(strings.Split(sc.Text(), sep), lvl)

		curr, next, fails := 0, 1, 0
		ord := order(lvl[curr], lvl[next])

		for next < len(lvl) {
			diff := pkg.AbsDiff(lvl[curr], lvl[next])

			if (diff >= 1 && diff <= 3) && (order(lvl[curr], lvl[next]) == ord) {
				curr = next
				next++
				continue
			}

			if fails >= tolerance {
				safe--
				break
			}

			fails++
			curr++
			next++
		}

		safe++
	}

	return safe
}

func order(a, b int) int {
	if a > b {
		return ordDesc
	}

	return ordAsc
}

func main() {
	part1 := solve(pkg.NewScanner(input), 0)
	part2 := solve(pkg.NewScanner(input), 1)

	fmt.Println(part1, part2)
}
