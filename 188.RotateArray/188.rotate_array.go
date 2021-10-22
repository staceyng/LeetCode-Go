package leetcode

import (
	"fmt"
)

// Given an array, rotate the array to the right by k steps, where k is non-negative.
// Rotate in place, find out why k%n is used instead of k to check when len(nums) != k, got stuck at k and test4
// This approach is based on the fact that when we rotate the array k times, k%nk elements from the back end of the array come to the front and the rest of the elements from the front shift backwards.
// Solution is better discussed here https://leetcode.com/problems/rotate-array/discuss/269948/4-solutions-in-python-(From-easy-to-hard)
// Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
// Could you do it in-place with O(1) extra space?

func rotate(nums []int, k int) []int {
	n := len(nums)
	k = k % n
	r := reverseInPlace(nums, 0, n-1)
	fmt.Println("reverse: ", r)
	rs := reverseInPlace(r, 0, k-1)
	fmt.Println("reverse_start: ", rs)
	re := reverseInPlace(rs, k, n-1)
	fmt.Println("reverse_end: ", re)
	return re
}

func reverseInPlace(nums []int, start int, end int) []int {
	left := start
	right := end

	for left <= right {
		nums[left], nums[right] = nums[right], nums[left]
		left += 1
		right -= 1
	}

	return nums
}
