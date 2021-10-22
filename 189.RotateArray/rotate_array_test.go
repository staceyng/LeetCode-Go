package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type inputParams struct {
	nums []int
	k    int
}

func TestRotateArray(t *testing.T) {
	var tests = []struct {
		name      string
		input     inputParams
		expOutput []int
	}{
		{
			"test 1",
			inputParams{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3},
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"test 2",
			inputParams{nums: []int{-1, -100, 3, 99}, k: 2},
			[]int{3, 99, -1, -100},
		},
		{
			"test 3",
			inputParams{nums: []int{-1}, k: 2},
			[]int{-1},
		},
		{
			"test 4",
			inputParams{nums: []int{1, 2}, k: 3},
			[]int{2, 1},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			actual := rotate(tt.input.nums, tt.input.k)

			assert.Equal(t, tt.expOutput, actual)
		})
	}
}
