package leetcode

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	res := make([]int, 2)
	for i < j {
		fmt.Println(numbers[i], numbers[j])
		if numbers[i]+numbers[j] > target {
			j -= 1
		} else if numbers[i]+numbers[j] < target {
			i += 1
		} else {
			res[0] = i + 1 // +1 for 1 indexed array
			res[1] = j + 1
			break
		}
	}
	return res
}
