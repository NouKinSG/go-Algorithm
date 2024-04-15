package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	return nil
}

func getKgroupEnd(start *ListNode, k int) *ListNode {
	for k > 1 && start != nil {
		start = start.Next
		k--
	}
	return start
}

func reverse(start, end *ListNode) {
	end = end.Next
	cur := start
	var pre *ListNode
	var next *ListNode

	for cur != end {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	start.Next = end
}


