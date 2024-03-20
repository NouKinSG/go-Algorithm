package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	ans := 0
	if root == nil {
		return ans
	}

	ans += pathFromNode1(root, targetSum)
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)

	return ans
}

// 给定节点为起点，往下 和为 targetSum的路径数
func pathFromNode1(node *TreeNode, Sum int) int {
	count := 0
	if node == nil {
		return count
	}

	if node.Val == Sum {
		count++
	}

	count += pathFromNode1(node.Left, Sum-node.Val)
	count += pathFromNode1(node.Right, Sum-node.Val)
	return count
}
