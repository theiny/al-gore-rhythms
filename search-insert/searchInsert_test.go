package searchinsert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nums     []int
	target   int
	expected int
}{
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3, 5, 6}, 7, 4},
	{[]int{1, 3, 5, 6}, 0, 0},
	{[]int{3, 6, 7, 8, 10}, 5, 1},
	{[]int{2, 3, 4, 7, 8, 9}, 11, 6},
}

func TestSearchInsert(t *testing.T) {
	for _, test := range tests {
		result := searchInsert(test.nums, test.target)
		assert.Equal(t, test.expected, result)
	}
}
