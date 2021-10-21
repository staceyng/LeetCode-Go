package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type inputParams struct {
	nums   []int
	target int
}

func TestSearchInsert(t *testing.T) {
	var tests = []struct {
		name      string
		input     inputParams
		expOutput int
	}{
		{
			"test 1",
			inputParams{nums: []int{1, 3, 5, 6}, target: 5},
			2,
		},
		{
			"test 2",
			inputParams{nums: []int{1, 3, 5, 6}, target: 2},
			1,
		},
		{
			"test 3",
			inputParams{nums: []int{1, 3, 5, 6}, target: 7},
			4,
		},
		{
			"test 4",
			inputParams{nums: []int{1, 3, 5, 6}, target: 0},
			0,
		},
		{
			"test 5",
			inputParams{nums: []int{1}, target: 0},
			0,
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			actual := searchInsert(tt.input.nums, tt.input.target)

			assert.Equal(t, tt.expOutput, actual)
		})
	}

}
