package parantheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  string
	output bool
}{
	{"{}", true},
	{"()[]{}", true},
	{"(]", false},
	{"([)]", false},
	{"{[]}", true},
	{"([]", false},
	{"[])", false},
	{"(([]){})", true},
	{"{{)}", false},
}

func TestIsValid(t *testing.T) {
	for _, test := range tests {
		result := isValid(test.input)
		assert.Equal(t, test.output, result)
	}
}
