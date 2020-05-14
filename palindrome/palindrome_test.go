package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  int
	output bool
}{
	{121, true},
	{-121, false},
	{10, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range tests {
		result := isPalindrome(test.input)
		assert.Equal(t, test.output, result)
	}
}
