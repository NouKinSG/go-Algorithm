package main

func isHuiWen(s string) bool {
	// 递归终止条件
	if len(s) <= 1 {
		return true
	}

	// 回文判断
	if s[0] != s[len(s)-1] {
		return false
	}
	return isHuiWen(s[1 : len(s)-1])
}

func contRepeat(s string) {

}

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// 创建一个栈用于存储节点
	stack := []*TreeNode{root}
	// 创建一个切片用于存储后序遍历的结果
	result := []int{}

	for len(stack) > 0 {
		// 弹出栈顶元素
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 将当前节点的值添加到结果切片中
		result = append(result, node.val)

		// 如果当前节点有右子树，将其压入栈中
		if node.right != nil {
			stack = append(stack, node.right)
		}

		if node.left != nil {
			stack = append(stack, node.left)
		}
	}
	return result
}
