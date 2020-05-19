package linkedduplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var l1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	},
}

func TestDeleteDuplicates(t *testing.T) {
	l := deleteDuplicates(l1)

	for l.Next != nil {
		assert.Equal(t, 1, l.Val)
		l = l.Next
		assert.Equal(t, 2, l.Val)
	}
}
