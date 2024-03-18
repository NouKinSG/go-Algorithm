package main

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func lowestCommonAncestor(root, p, q *TreeNode1) *TreeNode1 {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)   // 左子树查找
	right := lowestCommonAncestor(root.Right, p, q) // 右子树查找

	if left != nil && right != nil {
		// 如果 p 和 q 分别位于节点的两侧，那么当前节点即为它们的最近公共祖先
		return root
	}
	if left == nil {
		// 如果右子树找到了 p 或 q，返回右子树查找的结果
		return right
	}
	return left
}
