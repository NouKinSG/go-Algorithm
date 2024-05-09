package code06sortlisttobst

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

// sortedListToBST 将有序链表转换为高度平衡的二叉搜索树
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	midpre := findMidPre(head)
	if midpre == nil { // 当前链表只有两个节点，中点的前一个节点可能是nil
		return &TreeNode{Val: head.Val, Right: sortedListToBST(head.Next)}
	}
	mid := midpre.Next
	midpre.Next = nil // 断开链表
	root := &TreeNode{Val: mid.Val}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(mid.Next)
	return root
}

// 使用快慢指针找到链表中间节点的前一个节点
func findMidPre(head *ListNode) *ListNode {
	var pre *ListNode
	slow := head
	fast := head
	if head == nil || head.Next == nil { // 对边界情况的处理
		return nil
	}
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	return pre
}

// 快慢指针找中间节点
func findMidPtr(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 快慢指针找中间节点并且截断
func findMidPtrTruncation(head *ListNode) *ListNode {
	var pre *ListNode
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if pre != nil {
		pre.Next = nil
	}
	return slow
}
