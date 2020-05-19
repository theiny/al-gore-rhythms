package substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input    string
	expected int
}{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
	{" ", 1},
	{"dvdf", 3},
	{"brnk", 4},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, test := range tests {
		result := lengthOfLongestSubstring(test.input)
		assert.Equal(t, test.expected, result)
	}
}
