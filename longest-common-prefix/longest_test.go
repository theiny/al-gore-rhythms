package longest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  []string
	output string
}{
	{[]string{"flower", "flow", "flight"}, "fl"},
	{[]string{"dog", "racecar", "car"}, ""},
	{[]string{"a"}, "a"},
	{[]string{"a", "b"}, ""},
	{[]string{"c", "acc", "ccc"}, ""},
	{[]string{"aca", "cba"}, ""},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, test := range tests {
		result := longestCommonPrefix(test.input)
		assert.Equal(t, test.output, result)
	}
}
