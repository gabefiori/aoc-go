package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/gabefiori/aoc-go/pkg"
)

const (
	mulPrefix = "mul("
	maxMulLen = len("XXX,XXX)")

	doFn   = "do()"
	dontFn = "don't()"
)

//go:embed input.txt
var input string

func solve(s string, canDisable bool) int {
	pos, total := 0, 0
	isEnabled := true

	for pos+len(dontFn) < len(s) {
		if canDisable {
			if hasIdent(s, doFn, pos) {
				pos += len(doFn)
				isEnabled = true
			}

			if hasIdent(s, dontFn, pos) {
				pos += len(dontFn)
				isEnabled = false
			}
		}

		if !isEnabled || !hasIdent(s, mulPrefix, pos) {
			pos++
			continue
		}

		pos += len(mulPrefix)

		args := s[pos:pkg.Clamp(pos+maxMulLen, 0, len(s))]
		sep := strings.Index(args, ",")
		end := strings.Index(args, ")")

		if sep == -1 || end == -1 {
			pos++
			continue
		}

		total += pkg.ParseInt(args[:sep]) * pkg.ParseInt(args[sep+1:end])
		pos += end
	}

	return total
}

func hasIdent(s, ident string, pos int) bool {
	return s[pos:pos+len(ident)] == ident
}

func main() {
	part1 := solve(input, false)
	part2 := solve(input, true)

	fmt.Println(part1, part2)
}
