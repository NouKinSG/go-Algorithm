package tow

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归构建BST
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	midPrev := findMiddlePrev(head)
	if midPrev == nil { // 当链表只有两个节点时，中点的前一个节点可能是 nil
		return &TreeNode{Val: head.Val, Right: sortedListToBST(head.Next)}
	}
	mid := midPrev.Next
	midPrev.Next = nil // 断开链表

	root := &TreeNode{Val: mid.Val}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(mid.Next)
	return root
}

// 使用快慢指针找到链表中间节点的前一个节点
func findMiddlePrev(head *ListNode) *ListNode {
	var prevPtr *ListNode
	slowPtr := head
	fastPtr := head
	if head == nil || head.Next == nil { // 对于边界情况的处理
		return nil
	}
	for fastPtr != nil && fastPtr.Next != nil && fastPtr.Next.Next != nil {
		prevPtr = slowPtr
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}
	return prevPtr
}
