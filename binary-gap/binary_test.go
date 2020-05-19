package binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input    int
	expected int
}{
	{1041, 5}, //10000010001
	{15, 0},   //1111
	{32, 0},   //100000
}

func TestBinaryGap(t *testing.T) {
	for _, test := range tests {
		result := BinaryGap(test.input)
		assert.Equal(t, test.expected, result)
	}
}
