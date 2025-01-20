package four

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	// 自己不是 p 或者 q，找子树是不是
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 看看有没有找到
	if left != nil && right != nil {
		// 左右子树都找到了，说明是
		return root
	}

	//只找到任意一个，那就将找到的返回
	if left == nil {
		return right
	}
	return left
}
