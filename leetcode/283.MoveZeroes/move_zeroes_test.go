package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	var tests = []struct {
		name      string
		input     []int
		expOutput []int
	}{
		{
			"test 1",
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
		{
			"test 2",
			[]int{0},
			[]int{0},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			actual := moveZeroes(tt.input)

			assert.Equal(t, tt.expOutput, actual)
		})
	}
}
