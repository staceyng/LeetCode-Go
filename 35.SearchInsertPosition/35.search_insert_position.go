package leetcode

/* Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
You must write an algorithm with O(log n) runtime complexity. */

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) / 2)

		switch {
		case target == nums[mid]:
			return mid
		case target > nums[mid]:
			left = mid + 1
		case target < nums[mid]:
			right = mid - 1
		}
	}
	return left
}
