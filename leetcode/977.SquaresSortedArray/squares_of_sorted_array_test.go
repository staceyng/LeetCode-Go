package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	var tests = []struct {
		name      string
		input     []int
		expOutput []int
	}{
		{
			"test 1",
			[]int{-4, -1, 0, 3, 10},
			[]int{0, 1, 9, 16, 100},
		},
		{
			"test 2",
			[]int{-7, -3, 2, 3, 11},
			[]int{4, 9, 9, 49, 121},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			actual := sortedSquares(tt.input)

			assert.Equal(t, tt.expOutput, actual)
		})
	}

}
