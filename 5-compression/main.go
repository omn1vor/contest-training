package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var setCount int
	fmt.Fscanln(in, &setCount)

	for i := 0; i < setCount; i++ {
		var count int
		fmt.Fscanln(in, &count)

		nums := make([]int, count)
		for i := range nums {
			fmt.Scan(in, &nums[i])
		}

		res := []int{}
		count = 0
		direction := 1

		for i, num := range nums {
			if count == 0 {
				res = append(res, num)
			}

			descLen := scanLen(nums, i, -1)
			ascLen := scanLen(nums, i, 1)

		}

	}
}

func scanLen(nums []int, idx, direction int) int {
	count := 0
	for idx < len(nums)-1 && nums[idx] == nums[idx+1*direction] {
		count++
		idx++
	}
	return count
}
