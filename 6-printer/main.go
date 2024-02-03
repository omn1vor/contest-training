package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var setCount int
	fmt.Fscan(in, &setCount)

	for i := 0; i < setCount; i++ {
		fmt.Println(process(in))
	}
}

func process(in *bufio.Reader) string {
	var total int
	fmt.Fscan(in, &total)

	var input string
	fmt.Fscan(in, &input)

	printed := map[int]bool{}
	for _, new := range strings.Split(input, ",") {
		addPages(printed, new)
	}

	pages := []string{}
	buf := []int{}
	for i := 1; i <= total; i++ {
		_, found := printed[i]
		if !found {
			buf = append(buf, i)
		} else {
			pages = addBufferedPages(pages, buf)
			buf = buf[:0]
		}
	}
	pages = addBufferedPages(pages, buf)
	return strings.Join(pages, ",")
}

func addPages(printed map[int]bool, new string) {
	p := strings.Split(new, "-")
	start, _ := strconv.Atoi(p[0])
	end := start
	if len(p) > 1 {
		end, _ = strconv.Atoi(p[1])
	}
	for ; start <= end; start++ {
		printed[start] = true
	}
}

func addBufferedPages(pages []string, buf []int) []string {
	if len(buf) == 0 {
		return pages
	}

	if len(buf) == 1 {
		return append(pages, strconv.Itoa(buf[0]))
	} else {
		return append(pages, fmt.Sprintf("%d-%d", buf[0], buf[len(buf)-1]))
	}
}
