package six

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
	// 凑齐了一组
	head = end
	reverse(start, end)
	// 反转后，上一组结尾，是start
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

// 数 k 个返回
func getKGroupEnd(start *ListNode, k int) *ListNode {
	for k > 1 && start != nil {
		start = start.Next
		k--
	}
	return start
}

// 反转
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
