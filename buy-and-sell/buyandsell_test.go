package buyandsell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input    []int
	expected int
}{
	{[]int{7, 1, 5, 3, 6, 4}, 7},
	{[]int{1, 2, 3, 4, 5}, 4},
	{[]int{7, 6, 4, 3, 1}, 0},
}

func TestMaxProfit(t *testing.T) {
	for _, test := range tests {
		result := maxProfit(test.input)
		assert.Equal(t, test.expected, result)
	}
}
