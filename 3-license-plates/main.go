package main

import (
	"fmt"
	"os"
	"strings"
)

var templates = []string{
	"01100",
	"0100",
}

func main() {
	in := os.Stdin

	var count int
	fmt.Fscan(in, &count)

	var input string
	for i := 0; i < count; i++ {
		fmt.Fscan(in, &input)
		fmt.Println(validPlates(input))
	}

}

func validPlates(input string) string {
	plates := findPlates(input)
	if len(plates) == 0 {
		return "-"
	} else {
		return strings.Join(plates, " ")
	}
}

func toOnesZeroes(s string) string {
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= '0' && bytes[i] <= '9' {
			bytes[i] = '1'
		} else {
			bytes[i] = '0'
		}
	}
	return string(bytes)
}

func findPlates(s string) []string {
	input := toOnesZeroes(s)
	plates := []string{}
	count := 0

	for start := 0; start < len(input); {
		found := false
		for _, template := range templates {
			end := start + len(template)
			if end <= len(input) && template == input[start:end] {
				plates = append(plates, s[start:end])
				found = true
				count += end - start
				start = end
				break
			}
		}
		if !found {
			break
		}
	}

	if count != len(input) {
		return []string{}
	}
	return plates
}
