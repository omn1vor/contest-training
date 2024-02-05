package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type node struct {
	parent string
	id     string
	text   string
}

func main() {
	in := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(in)
	scanner.Scan()
	var setCount int
	fmt.Sscan(scanner.Text(), &setCount)

	for i := 0; i < setCount; i++ {
		process(scanner)
	}
}

func process(scanner *bufio.Scanner) {
	scanner.Scan()
	var count int
	fmt.Sscan(scanner.Text(), &count)

	nodes := map[string][]node{}

	for i := 0; i < count; i++ {
		scanner.Scan()
		text := scanner.Text()
		id, text, _ := strings.Cut(text, " ")
		parent, text, _ := strings.Cut(text, " ")
		nodes[parent] = append(nodes[parent], node{parent, id, text})
	}

	for _, children := range nodes {
		sort.Slice(children, func(i, j int) bool { return children[i].id < children[j].id })
	}

	branch := nodes["-1"]
	printBranch(branch, nodes, "")
}

func printBranch(b []node, nodes map[string][]node, prefix string) {
	for i, node := range b {
		if i > 0 && prefix == "" {
			fmt.Println()
		}
		if prefix != "" {
			fmt.Println(prefix)
		}
		fmt.Print(prefix)
		if prefix != "" {
			fmt.Print("--")
		}
		fmt.Println(node.text)
		hasSiblingDown := len(b) > i+1
		newPrefix := "|"
		if prefix != "" {
			newPrefix = "  " + newPrefix
		}

		if !hasSiblingDown {
			if len(prefix) > 1 {
				newPrefix = prefix[:len(prefix)-3] + "   " + newPrefix
			} else {
				newPrefix = "" + newPrefix
			}
		} else {
			newPrefix = prefix + newPrefix
		}
		printBranch(nodes[node.id], nodes, newPrefix)
	}
}
