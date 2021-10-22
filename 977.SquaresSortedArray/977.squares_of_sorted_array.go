package leetcode

import (
	"fmt"
)

// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
// find a way to achieve O(n) runtime vs square array followed by sort

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	res := make([]int, len(nums))
	idx := len(nums) - 1
	// fmt.Println(left, right, idx)

	for left <= right {
		leftSq := square(nums[left])
		rightSq := square(nums[right])

		if leftSq < rightSq {
			res[idx] = rightSq
			right -= 1
			// fmt.Println(res)
		} else {
			res[idx] = leftSq
			left += 1
			// fmt.Println(res)
		}
		idx -= 1
	}

	return res

}

func square(n int) int {
	return n * n
}
