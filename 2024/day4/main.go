package main

import (
	"fmt"
	"strings"

	_ "embed"
)

var dirs = [...][2]int{
	{1, 0},   // right
	{-1, 0},  // left
	{0, -1},  // up
	{0, 1},   // down
	{-1, -1}, // up-left (diagonal)
	{1, -1},  // up-right (diagonal)
	{-1, 1},  // down-left (diagonal)
	{1, 1},   // down-right (diagonal)
}

//go:embed input.txt
var input string

func solve(s string) (int, int) {
	var xmas, mas int
	lines := strings.Split(strings.TrimSpace(s), "\n")

	for i, line := range lines {
		for j := range line {
			if checkMas(lines, j, i) {
				mas++
			}

			for _, dir := range dirs {
				if checkXmas(lines, j, i, dir[0], dir[1]) {
					xmas++
				}
			}
		}
	}

	return xmas, mas
}

func checkXmas(lines []string, x, y, dx, dy int) bool {
	if lines[y][x] != 'X' {
		return false
	}

	lx := x + (dx * 3)
	ly := y + (dy * 3)

	if (ly < 0 || ly >= len(lines)) || (lx < 0 || lx >= len(lines[ly])) {
		return false
	}

	return lines[y+dy][x+dx] == 'M' &&
		lines[y+(dy*2)][x+(dx*2)] == 'A' &&
		lines[ly][lx] == 'S'
}

func checkMas(lines []string, x, y int) bool {
	if lines[y][x] != 'A' ||
		(y-1 < 0 || y+1 >= len(lines)) ||
		(x-1 < 0 || x+1 >= len(lines[y])) {
		return false
	}

	return lines[y-1][x-1] == 'M' && lines[y-1][x+1] == 'M' && lines[y+1][x-1] == 'S' && lines[y+1][x+1] == 'S' ||
		lines[y+1][x-1] == 'M' && lines[y+1][x+1] == 'M' && lines[y-1][x-1] == 'S' && lines[y-1][x+1] == 'S' ||
		lines[y-1][x-1] == 'M' && lines[y+1][x-1] == 'M' && lines[y-1][x+1] == 'S' && lines[y+1][x+1] == 'S' ||
		lines[y-1][x+1] == 'M' && lines[y+1][x+1] == 'M' && lines[y-1][x-1] == 'S' && lines[y+1][x-1] == 'S'
}

func main() {
	part1, part2 := solve(input)
	fmt.Println(part1, part2)
}
