package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  string
	output int
}{
	{"III", 3},
	{"IV", 4},
	{"IX", 9},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
}

func TestRomanToInt(t *testing.T) {
	for _, test := range tests {
		result := romanToInt(test.input)
		assert.Equal(t, test.output, result)
	}
}
