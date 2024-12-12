package main

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"

	"github.com/gabefiori/aoc-go/pkg"
)

const (
	ordInit int = iota
	ordAsc
	ordDesc
)

//go:embed input.txt
var input []byte

func solve(sc *bufio.Scanner, tolerate bool) int {
	var count int

	lvl := make([]int, 10)
	newLvl := make([]int, 10)

	for sc.Scan() {
		lvl := pkg.ParseInts(strings.Split(sc.Text(), " "), lvl)

		if safe(lvl) {
			count++
		} else if tolerate {
			for i := 0; i < len(lvl); i++ {
				newLvl = newLvl[:0]
				newLvl = append(newLvl, lvl[:i]...)
				newLvl = append(newLvl, lvl[i+1:]...)

				if safe(newLvl) {
					count++
					break
				}
			}
		}
	}

	return count
}

func safe(lvl []int) bool {
	curr, next, ord := 0, 1, ordInit

	for next < len(lvl) {
		diff := pkg.AbsDiff(lvl[curr], lvl[next])

		if diff < 1 || diff > 3 {
			return false
		}

		newOrd := order(lvl[curr], lvl[next])
		if ord != ordInit && ord != newOrd {
			return false
		}

		ord = newOrd
		curr = next
		next++
	}

	return true
}

func order(a, b int) int {
	if a > b {
		return ordDesc
	}

	if a < b {
		return ordAsc
	}

	return ordInit
}

func main() {
	part1 := solve(pkg.NewScanner(input), false)
	part2 := solve(pkg.NewScanner(input), true)
	fmt.Println(part1, part2)
}
