package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"maps"
	"slices"

	"github.com/gabefiori/aoc-go/pkg"
)

//go:embed input.txt
var input []byte

type position struct {
	x, y int
}

func solve(sc *bufio.Scanner) (int, int) {
	var start, dir position
	var maze [][]byte

	for row := 0; sc.Scan(); row++ {
		line := sc.Text()
		maze = append(maze, append([]byte(nil), line...))

		for col, b := range line {
			if b == '^' {
				dir.y = -1
				start = position{col, row}
			}
		}
	}

	_, visited := walk(maze, start, dir)
	loops := 0

	for _, v := range visited {
		if maze[v.y][v.x] == '^' {
			continue
		}

		maze[v.y][v.x] = '#'
		if isLoop, _ := walk(maze, start, dir); isLoop {
			loops++
		}
		maze[v.y][v.x] = '.'
	}

	return len(visited), loops
}

func walk(maze [][]byte, pos position, dir position) (bool, []position) {
	visited := make(map[position][]position, 6000)

	for {
		if (pos.x < 0 || pos.x >= len(maze[0])) || (pos.y < 0 || pos.y >= len(maze)) {
			return false, slices.Collect(maps.Keys(visited))
		}

		if d, seen := visited[pos]; seen && slices.Index(d, dir) >= 0 {
			return true, nil
		}

		if maze[pos.y][pos.x] == '#' {
			pos.x -= dir.x
			pos.y -= dir.y

			dir.x, dir.y = -dir.y, dir.x
		}

		visited[pos] = append(visited[pos], dir)

		pos.x += dir.x
		pos.y += dir.y
	}
}

func main() {
	part1, part2 := solve(pkg.NewScanner(input))
	fmt.Println(part1, part2)
}
