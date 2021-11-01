package leetcode

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	res := []string{}
	resultString := ""
	for _, v := range ss {
		fmt.Println(v)
		r := reverse(v)
		fmt.Println(r)
		res = append(res, r)
	}

	for i := 0; i <= len(res)-1; i++ {
		resultString += res[i] + " "
	}

	return strings.TrimRight(resultString, " ")

}

func reverse(s string) string {
	b := []byte(s)
	left := 0
	right := len(b) - 1

	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}

	return string(b)
}
