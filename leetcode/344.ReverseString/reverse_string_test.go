package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	var tests = []struct {
		name      string
		input     []byte
		expOutput []byte
	}{
		{
			"test 1",
			[]byte("hello"),
			[]byte("olleh"),
		},
		{
			"test 2",
			[]byte("Hannah"),
			[]byte("hannaH"),
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			actual := reverseString(tt.input)

			assert.Equal(t, tt.expOutput, actual)
		})
	}
}
