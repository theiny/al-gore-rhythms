package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	haystack string
	needle   string
	expected int
}{
	{"hello", "ll", 2},
	{"aaaaa", "bba", -1},
}

func TestStrStr(t *testing.T) {
	for _, test := range tests {
		result := strStr(test.haystack, test.needle)
		assert.Equal(t, test.expected, result)
	}
}
