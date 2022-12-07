package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getArray(l *ListNode) (r []int) {
	for l != nil {
		r = append(r, l.Val)
        l = l.Next
	}
	return
}

var (
	l1 = ListNode{Val: 3, Next: nil}
	l2 = ListNode{Val: 4, Next: &l1}
	l3 = ListNode{Val: 2, Next: &l2}
)

var (
	r1 = ListNode{Val: 4, Next: nil}
	r2 = ListNode{Val: 6, Next: &r1}
	r3 = ListNode{Val: 5, Next: &r2}
)

func TestAddTwoNumbers(t *testing.T) {
	assert.Equal(t, []int{7, 0, 8}, getArray(addTwoNumbers(&l3, &r3)))
}
