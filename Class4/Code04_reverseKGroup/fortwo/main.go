package main

type ListNode struct {
	val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	end := getKGroupEnd(start, k)
	if end == nil {
		return head
	}
	// 第一组凑齐了
	head = end
	reverse(start, end)
	// 反转后，上一组的结尾节点，是start
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

// 给一个头节点，数 k 个，返回头节点
func getKGroupEnd(start *ListNode, k int) *ListNode {

	for k > 1 && start != nil {
		start = start.Next
		k--
	}
	return start
}

// 反转某组链表
func reverse(start, end *ListNode) {
	end = end.Next
	var pre *ListNode
	cur := start
	var next *ListNode

	for cur != end {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	start.Next = end
}
