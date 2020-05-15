package duplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input    []int
	expected int
}{
	{[]int{1, 1, 2}, 2},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, test := range tests {
		length := removeDuplicates(test.input)
		assert.Equal(t, test.expected, length)
	}
}
