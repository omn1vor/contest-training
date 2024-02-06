package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	parent int
	id     int
	text   string
}

func main() {
	in := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(in)
	scanner.Scan()
	var setCount int
	fmt.Sscan(scanner.Text(), &setCount)

	for i := 0; i < setCount; i++ {
		fmt.Println(process(scanner))
	}
}

func process(scanner *bufio.Scanner) string {
	scanner.Scan()
	var count int
	fmt.Sscan(scanner.Text(), &count)

	nodes := map[int][]node{}

	for i := 0; i < count; i++ {
		scanner.Scan()
		text := scanner.Text()
		idStr, text, _ := strings.Cut(text, " ")
		parentStr, text, _ := strings.Cut(text, " ")
		id, _ := strconv.Atoi(idStr)
		parent, _ := strconv.Atoi(parentStr)
		nodes[parent] = append(nodes[parent], node{parent, id, text})
	}

	for _, children := range nodes {
		sort.Slice(children, func(i, j int) bool { return children[i].id < children[j].id })
	}

	branch := nodes[-1]
	return printBranch(branch, nodes, "")
}

func printBranch(branch []node, nodes map[int][]node, prefix string) string {
	b := &strings.Builder{}
	for i, node := range branch {
		if i > 0 && prefix == "" {
			fmt.Fprintln(b)
		}
		if prefix != "" {
			fmt.Fprintln(b, prefix)
		}
		fmt.Fprint(b, prefix)
		if prefix != "" {
			fmt.Fprint(b, "--")
		}
		fmt.Fprintln(b, node.text)
		hasSiblingDown := len(branch) > i+1
		newPrefix := "|"
		if prefix != "" {
			newPrefix = "  |"
		}
		if !hasSiblingDown && prefix != "" {
			if prefix == "|" {
				newPrefix = " " + newPrefix
			} else {
				newPrefix = prefix[:len(prefix)-3] + "   " + newPrefix
			}
		} else {
			newPrefix = prefix + newPrefix
		}
		b.WriteString(printBranch(nodes[node.id], nodes, newPrefix))
	}
	return b.String()
}
