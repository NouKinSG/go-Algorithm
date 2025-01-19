package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)   // 找左
	right := lowestCommonAncestor(root.Right, p, q) // 找右

	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
