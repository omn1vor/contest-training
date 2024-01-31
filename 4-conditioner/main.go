package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	lt = "<="
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var count int
	fmt.Fscanln(in, &count)

	for i := 0; i < count; i++ {
		var empCount int
		fmt.Fscanln(in, &empCount)
		minTemp := 15
		maxTemp := 30
		var operand, temperature string
		for j := 0; j < empCount; j++ {
			fmt.Fscanln(in, &operand, &temperature)
			temp, _ := strconv.Atoi(temperature)
			if operand == lt {
				if maxTemp > temp {
					maxTemp = temp
				}
			} else {
				if minTemp < temp {
					minTemp = temp
				}
			}
			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println("-1")
			}
		}
		fmt.Println()
	}
}
