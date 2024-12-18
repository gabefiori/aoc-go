package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"maps"

	"github.com/gabefiori/aoc-go/pkg"
)

type position struct {
	x, y int
}

//go:embed input.txt
var input []byte
var dirs = []position{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func solve(sc *bufio.Scanner) (int, int) {
	tMap := make([][]byte, 0)
	heads := make([]position, 0, 180)

	for row := 0; sc.Scan(); row++ {
		line := sc.Bytes()
		tMap = append(tMap, append([]byte(nil), line...))

		for col, b := range line {
			if b == '0' {
				heads = append(heads, position{col, row})
			}
		}
	}

	seens := make([]map[position]int, len(heads))
	scores, ratings := 0, 0
	for i, h := range heads {
		seens[i] = make(map[position]int, 10)
		scores += walk(tMap, h, seens[i], tMap[h.y][h.x], 0)

		for r := range maps.Values(seens[i]) {
			ratings += r
		}
	}

	return scores, ratings
}

func walk(trail [][]byte, pos position, seen map[position]int, curr byte, count int) int {
	if curr == '9' {
		if _, exists := seen[pos]; !exists {
			count++
		}

		seen[pos]++
		return count
	}

	for _, dir := range dirs {
		x, y := pos.x+dir.x, pos.y+dir.y

		if x < 0 || x >= len(trail[0]) || y < 0 || y >= len(trail) {
			continue
		}

		if trail[y][x] == curr+1 {
			count = walk(trail, position{x, y}, seen, curr+1, count)
		}
	}

	return count
}

func main() {
	part1, part2 := solve(pkg.NewScanner(input))
	fmt.Println(part1, part2)
}
