package leetcode

import (
	"fmt"
)

// visual solution https://leetcode.com/problems/move-zeroes/discuss/979820/Two-Pointeror-Visual-or-Python

func moveZeroes(nums []int) []int {
	i := 0
	for j := range nums {
		fmt.Println(nums[i], nums[j])
		if nums[j] != 0 {
			// fmt.Println("in loop", i, j)
			nums[i], nums[j] = nums[j], nums[i]
			i += 1
		}

	}
	return nums
}
