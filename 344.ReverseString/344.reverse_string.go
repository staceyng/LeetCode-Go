package leetcode

import "fmt"

func reverseString(s []byte) []byte {
	// iterative
	left := 0
	right := len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left += 1
		right -= 1
	}

	fmt.Println(s)

	return s
}
