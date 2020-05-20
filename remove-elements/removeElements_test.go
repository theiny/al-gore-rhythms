package remove

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nums     []int
	value    int
	expected int
}{
	{[]int{3, 2, 2, 3}, 3, 2},
	{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
}

func TestRemoveElement(t *testing.T) {
	for _, test := range tests {
		result := removeElement(test.nums, test.value)
		assert.Equal(t, test.expected, result)
	}
}
