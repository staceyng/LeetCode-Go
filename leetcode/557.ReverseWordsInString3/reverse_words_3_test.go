package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	var tests = []struct {
		name      string
		input     string
		expOutput string
	}{
		{
			"test 1",
			"Let's take LeetCode contest",
			"s'teL ekat edoCteeL tsetnoc",
		},
		{
			"test 2",
			"God Ding",
			"doG gniD",
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			actual := reverseWords(tt.input)

			assert.Equal(t, tt.expOutput, actual)
		})
	}
}
