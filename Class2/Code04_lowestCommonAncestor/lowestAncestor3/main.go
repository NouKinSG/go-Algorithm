package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	// 自己不是 p 或者 q，找子树看看是不是
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果 left 和 right 都不为空 说明，找到了 p 和 q ，并且是自己的子树
	// 那么 root 就是 最近公共祖先
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}

// Define your data structures and variables here

func main() {
	// Your code logic goes here
}
