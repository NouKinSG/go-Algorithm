package three

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

// findMiddlePre 寻找链表的中间节点的前一个节点
func findMiddlePre(head *ListNode) *ListNode {
	var pre *ListNode
	slow := head
	fast := head
	if head == nil || head.Next == nil {
		return nil
	}
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	return pre
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	midPre := findMiddlePre(head)
	if midPre == nil {
		return &TreeNode{Val: head.Val, Right: sortedListToBST(head.Next)}
	}
	mid := midPre.Next
	// 断开链表
	midPre.Next = nil

	// 构造二叉树
	root := &TreeNode{Val: mid.Val}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(mid.Next)
	return root
}
