package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var setCount int
	fmt.Fscan(in, &setCount)

	for i := 0; i < setCount; i++ {
		compressed := compress(in)
		fmt.Println(len(compressed))
		strs := make([]string, len(compressed))
		for i, num := range compressed {
			strs[i] = strconv.Itoa(num)
		}
		fmt.Println(strings.Join(strs, " "))

	}
}

func compress(in *bufio.Reader) []int {
	var count int
	fmt.Fscan(in, &count)

	nums := make([]int, count)
	for i := range nums {
		_, err := fmt.Fscan(in, &nums[i])
		if err != nil {
			panic(err)
		}
	}

	res := []int{}
	count = 0
	last := math.MinInt
	direction := 0

	for i, num := range nums {
		possibleDirection := getDirection(last, num)
		if possibleDirection == direction && possibleDirection != 0 {
			count++
		} else {
			if i != 0 {
				res = append(res, count*direction)
			}
			count = 0
			direction = 0
			if i+1 < len(nums) {
				direction = getDirection(num, nums[i+1])
			}
			res = append(res, num)
		}
		last = num
	}
	res = append(res, count*direction)
	return res
}

func getDirection(prev, next int) int {
	if next == prev+1 {
		return 1
	} else if next == prev-1 {
		return -1
	}
	return 0
}
