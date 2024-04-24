package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	count := 0
	for cur != nil {
		cur = cur.Next
		count++
	}

	m := count - n
	if m == 0 {
		return head.Next
	}
	cur = head
	for i := 1; i <= m; i++ {
		// 找到n的前一个节点
		if i == m {
			cur.Next = cur.Next.Next
			return cur
		}
		cur = cur.Next
	}
	return head
}
