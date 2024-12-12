package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"

	_ "embed"

	"github.com/gabefiori/aoc-go/pkg"
)

const sep = "   "

//go:embed input.txt
var input []byte

func solve(sc *bufio.Scanner) (int, int) {
	var left, right []int
	counts := make(map[int]int)

	for sc.Scan() {
		str := sc.Text()
		idx := strings.Index(str, sep)

		r := pkg.ParseInt(str[idx+len(sep):])

		left = append(left, pkg.ParseInt(str[0:idx]))
		right = append(right, r)

		counts[r]++
	}

	slices.Sort(left)
	slices.Sort(right)

	sum, sim := 0, 0

	for i := 0; i < len(left); i++ {
		l := left[i]

		sum += pkg.AbsDiff(l, right[i])

		if c, exists := counts[l]; exists {
			sim += l * c
		}
	}

	return sum, sim
}

func main() {
	fmt.Println(pkg.NewScanner(input))
}
