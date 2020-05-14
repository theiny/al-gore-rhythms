package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nums     []int
	target   int
	expected []int
}{
	{
		[]int{2, 7, 11, 15},
		9,
		[]int{0, 1},
	},
	{
		[]int{1, 3, 4, 2},
		6,
		[]int{2, 3},
	},
	{
		[]int{3, 3},
		6,
		[]int{0, 1},
	},
	{
		[]int{3, 2, 4},
		6,
		[]int{1, 2},
	},
}

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		r := twoSum(test.nums, test.target)
		assert.ElementsMatch(t, test.expected, r)
	}
}
