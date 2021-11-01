package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		name     string
		input    int
		expOuput int
	}{
		{
			"test 1 - element exists in array",
			5,
			4,
		},
		{
			"test 2 - element does not exist in array",
			1,
			1,
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			actual := firstBadVersion(tt.input)

			assert.Equal(t, tt.expOuput, actual)
		})
	}
}
