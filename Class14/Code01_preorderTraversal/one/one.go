package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144.二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		// 子节点入栈
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return ans
}
