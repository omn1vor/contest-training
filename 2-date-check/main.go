package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	shortMax   = 30
	febMax     = 28
	febMaxleap = 29
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		data := strings.Split(line, " ")
		if len(data) == 1 {
			continue
		}
		fmt.Println(boolString(checkDate(data)))
	}
}

func checkDate(dates []string) bool {
	day, _ := strconv.Atoi(dates[0])
	month, _ := strconv.Atoi(dates[1])
	year, _ := strconv.Atoi(dates[2])

	if month == 2 {
		if isLeap(year) {
			return day <= febMaxleap
		}
		return day <= febMax
	} else {
		isShort := month < 8 && month%2 == 0 || month > 7 && month%2 != 0
		if isShort {
			return day <= shortMax
		}
		return true
	}
}

func isLeap(year int) bool {
	return year%400 == 0 || year%100 != 0 && year%4 == 0
}

func boolString(val bool) string {
	if val {
		return "YES"
	}
	return "NO"
}
