package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	end := getKGroupEnd(start, k)
	if end == nil {
		return head
	}

	head = end
	reverse(start, end)
	lastEnd := start
	for lastEnd.Next != nil {
		start = lastEnd.Next
		end = getKGroupEnd(start, k)
		if end == nil {
			return head
		}
		reverse(start, end)
		lastEnd.Next = end
		lastEnd = start
	}
	return head
}

// 数 k 个返回。
func getKGroupEnd(start *ListNode, k int) *ListNode {
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
