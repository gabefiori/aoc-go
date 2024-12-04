package pkg

import (
	"iter"
)

func ParseIntsIter(ss []string) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, s := range ss {
			if !yield(ParseInt(s)) {
				return
			}
		}
	}
}

func ParseInts(s []string, pre ...[]int) []int {
	var t []int

	if pre == nil {
		t = make([]int, 0, len(s))
	} else {
		t = pre[0][:0]
	}

	for v := range ParseIntsIter(s) {
		t = append(t, v)
	}

	return t
}
