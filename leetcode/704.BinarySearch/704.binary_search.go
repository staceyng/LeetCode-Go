package leetcode

func search(nums []int, target int) int {
	lp, rp := 0, len(nums)-1

	for lp <= rp {
		mp := lp + ((rp - lp) / 2)

		switch {
		case nums[mp] == target:
			return mp
		case target < nums[mp]:
			rp = mp - 1
		case target > nums[mp]:
			lp = mp + 1
		}

	}
	return -1
}
