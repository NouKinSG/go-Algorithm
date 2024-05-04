package two

type ListNode struct {
	Val  int
	Next *ListNode
}

// 19. 删除链表的倒数第 N 个结点
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
	cur = head
	if m == 0 {
		return head.Next
	}
	for i := 1; i <= m; i++ {
		if i == m {
			cur.Next = cur.Next.Next
			return head
		}
		cur = cur.Next
	}
	return head
}
