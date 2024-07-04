package main

import (
	"fmt"
)

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrderTraversal 使用栈实现前序遍历
func PreOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) > 0 {
		// 弹出栈顶节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 输出节点值
		fmt.Print(node.Val, " ")

		// 先压入右子树，再压入左子树
		// 这样保证左子树先遍历
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

func main() {
	// 构建一个简单的二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	// 前序遍历
	PreOrderTraversal(root)
}
