package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type inputParams704 struct {
	nums   []int
	target int
}

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		name     string
		input    inputParams704
		expOuput int
	}{
		{
			"test 1 - element exists in array",
			inputParams704{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9},
			4,
		},
		{
			"test 2 - element does not exist in array",
			inputParams704{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2},
			-1,
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			actual := search(tt.input.nums, tt.input.target)

			assert.Equal(t, tt.expOuput, actual)
		})
	}
}
