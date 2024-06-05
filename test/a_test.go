package test

import (
	"fmt"
	"testing"
)

// 链表反转
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func TestReverse(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}

	ans := reverse(head)

	fmt.Println(ans.Val)
	fmt.Println(ans.Next.Val)
	fmt.Println(ans.Next.Next.Val)
}
