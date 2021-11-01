package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type inputParams167 struct {
	nums   []int
	target int
}

func TestTwoSum2(t *testing.T) {
	var tests = []struct {
		name      string
		input     inputParams167
		expOutput []int
	}{
		{
			"test 1",
			inputParams167{nums: []int{2, 7, 11, 15}, target: 9},
			[]int{1, 2},
		},
		{
			"test 2",
			inputParams167{nums: []int{2, 3, 4}, target: 6},
			[]int{1, 3},
		},
		{
			"test 3",
			inputParams167{nums: []int{-1, 0}, target: -1},
			[]int{1, 2},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			actual := twoSum(tt.input.nums, tt.input.target)

			assert.Equal(t, tt.expOutput, actual)
		})
	}
}
