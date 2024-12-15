package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"maps"

	"github.com/gabefiori/aoc-go/pkg"
)

//go:embed input.txt
var input []byte

type position struct {
	x, y int
}

func solve(sc *bufio.Scanner, anyDist bool) int {
	antennas := make(map[rune][]position, 50)
	antinodes := make(map[position]struct{}, 1200)

	rows, cols := 0, 0
	for sc.Scan() {
		cols = len(sc.Text())
		for i, r := range sc.Text() {
			if r != '.' {
				antennas[r] = append(antennas[r], position{i, rows})
			}
		}

		rows++
	}

	dist := 1
	if anyDist {
		dist = cols
	}

	addAntinode := func(p position) {
		if p.x >= 0 && p.x < cols && p.y >= 0 && p.y < rows {
			antinodes[p] = struct{}{}
		}
	}

	for ant := range maps.Values(antennas) {
		for i := 0; i < len(ant)-1; i++ {
			for j := i + 1; j < len(ant); j++ {
				for k := 1; k <= dist; k++ {
					dX, dY := (ant[j].x-ant[i].x)*k, (ant[j].y-ant[i].y)*k

					addAntinode(position{ant[i].x - dX, ant[i].y - dY})
					addAntinode(position{ant[j].x + dX, ant[j].y + dY})

					if anyDist {
						antinodes[ant[i]] = struct{}{}
						antinodes[ant[j]] = struct{}{}
					}
				}
			}
		}
	}

	return len(antinodes)
}

func main() {
	part1 := solve(pkg.NewScanner(input), false)
	part2 := solve(pkg.NewScanner(input), true)
	fmt.Println(part1, part2)
}
