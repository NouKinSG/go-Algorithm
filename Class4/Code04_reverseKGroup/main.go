package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail := tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		next := tail.Next
		head, tail := myreverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}

	return hair.Next
}

func myreverse(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	p := head
	for pre != tail {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
	}
	return tail, head
}
