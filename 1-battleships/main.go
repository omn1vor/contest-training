package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	yes = "YES"
	no  = "NO"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ships := strings.Split(scanner.Text(), " ")
		if len(ships) == 1 {
			continue
		}
		fmt.Println(checkShips(ships))
	}
}

func checkShips(ships []string) string {
	counts := map[string]int{
		"1": 4,
		"2": 3,
		"3": 2,
		"4": 1,
	}

	for _, ship := range ships {
		ship = strings.TrimSpace(ship)
		counts[ship]--
		if counts[ship] < 0 {
			return no
		}
	}

	for _, v := range counts {
		if v != 0 {
			return no
		}
	}

	return yes
}
