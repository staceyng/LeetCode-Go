package leetcode

import "github.com/staceyng/LeetCode-Go/structures"
import "github.com/stretchr/testify/assert"
import "testing"

type inputParams19 struct {
	input []int
	n     int
}

func TestRemoveNthFromEnd(t *testing.T) {
	var tests = []struct {
		name      string
		input     inputParams19
		expOutput []int
	}{
		{
			"test 1",
			inputParams19{input: []int{1, 2, 3, 4, 5}, n: 2},
			[]int{1, 2, 3, 5},
		},
		{
			"test 2",
			inputParams19{input: []int{1}, n: 1},
			[]int{},
		},
		{
			"test 3",
			inputParams19{input: []int{1, 2}, n: 1},
			[]int{1},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			actualNode := removeNthFromEnd(structures.IntArrToLinkedList(tt.input.input), tt.input.n)
			actual := structures.LinkedListToIntArr(actualNode)

			assert.Equal(t, tt.expOutput, actual)
		})
	}
}
