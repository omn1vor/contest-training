package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	filled = '*'
	empty  = '.'
)

type rect struct {
	start, end, nestCount int
}

func newRect(start, end int) *rect {
	return &rect{start, end, 0}
}

type shapes struct {
	data []*rect
	idx  int
}

func newShapes() *shapes {
	return &shapes{[]*rect{}, 0}
}

func (s *shapes) push(r *rect) {
	r.nestCount = s.getNesting(r)
	s.data = append(s.data, r)
	s.idx++
}

func (s *shapes) finish(col int) (int, bool) {
	for i := len(s.data) - 1; i >= 0; i-- {
		if s.data[i].start == col {
			r := s.data[i]
			s.data = append(s.data[:i], s.data[i+1:]...)
			return r.nestCount, true
		}
	}
	return 0, false
}

func (s *shapes) getNesting(r *rect) int {
	var count int
	for _, shape := range s.data {
		if shape.start < r.start && shape.end > r.end {
			count++
		}
	}
	return count
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var setCount int
	fmt.Fscan(in, &setCount)

	for i := 0; i < setCount; i++ {
		process(in)
	}
}

func process(in *bufio.Reader) {
	var lines, cols int
	fmt.Fscan(in, &lines, &cols)

	shapes := newShapes()
	results := []int{}

	for line := 0; line < lines; line++ {
		var s string
		fmt.Fscan(in, &s)
		count := 0
		for col := 0; col < cols; col++ {
			if s[col] == filled {
				count++
			} else {
				if count > 1 {
					startOrEndShape(count, col, shapes, &results)
				}
				count = 0
			}
		}
		if count > 1 {
			startOrEndShape(count, cols, shapes, &results)
		}
	}

	sort.Slice(results, func(i, j int) bool { return results[i] < results[j] })
	fmt.Println(strings.Trim(fmt.Sprint(results), "[]"))
}

func startOrEndShape(count, col int, shapes *shapes, results *[]int) {
	start := col - count
	end := col - 1
	if surrounds, found := shapes.finish(start); found {
		*results = append(*results, surrounds)
	} else {
		shapes.push(newRect(start, end))
	}
}
