package reverse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	num    int
	expect int
}{
	{123, 321},
	{-123, -321},
	{1534236469, 0},
	{-2147483648, 0},
}

func TestReverse(t *testing.T) {
	for _, test := range tests {
		r := reverse(test.num)
		assert.Equal(t, test.expect, r)
	}
}

func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, test := range tests {
			b.Run(fmt.Sprintf("%d %d", test.num, test.expect), func(t *testing.B) {
				r := reverse(test.num)
				assert.Equal(b, test.expect, r)
			})
		}
	}
}
