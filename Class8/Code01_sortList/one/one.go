package one

//148. 排序链表
// Step 1: Split the list into two halves.
// Step 2: Recursively sort each half.
// Step 3: Merge the sorted halves.

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head *ListNode, tail *ListNode) *ListNode {
	if head == tail {
		return head
	}
	if head.Next == tail {
		return head
	}

	// Step 1: Split the list into two halves.
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != tail {
			fast = fast.Next
		}
	}
	mid := slow
	// Step 2: Recursively sort each half.
	return merge(sort(head, mid), sort(mid, tail))
}

// Step 3: Merge the sorted halves.
func merge(head1, head2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head2
			head1 = head1.Next
		} else {
			cur.Next = head1
			head2 = head2.Next
		}
		cur = cur.Next
	}
	if head1 != nil {
		cur.Next = head1
	}
	if head2 != nil {
		cur.Next = head2
	}
	return res.Next
}
