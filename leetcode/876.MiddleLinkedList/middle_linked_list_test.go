package leetcode

import (
	"github.com/staceyng/LeetCode-Go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddleLinkedList(t *testing.T) {
	var tests = []struct {
		name      string
		input     []int
		expOutput []int
	}{
		{
			"test 1",
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5},
		}, {
			"test 2",
			[]int{1, 2, 3, 4, 5, 6},
			[]int{4, 5, 6},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			actualNode := middleNode(structures.IntArrToLinkedList(tt.input))
			actual := structures.LinkedListToIntArr(actualNode)

			assert.Equal(t, tt.expOutput, actual)
		})
	}
}
