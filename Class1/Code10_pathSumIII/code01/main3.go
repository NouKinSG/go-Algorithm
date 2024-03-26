package code01

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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

	// 以root为起点出发，看看有没有 路径和为 targetSum的路径
	ans += pathFromNode(root, targetSum)
	ans += pathFromNode(root.Left, targetSum)
	ans += pathFromNode(root.Right, targetSum)
	return ans
}

// 从给定节点作为起点
func pathFromNode(node *TreeNode, targetSum int) int {
	count := 0
	if node == nil {
		return count
	}
	if node.Val == targetSum {
		count++
	}
	// 再看看 左子树和右子树有没有
	count += pathFromNode(node.Left, targetSum-node.Val)
	count += pathFromNode(node.Right, targetSum-node.Val)
	return count
}
