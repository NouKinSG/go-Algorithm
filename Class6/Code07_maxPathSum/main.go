package main

import "math"

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 124. 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选择对应子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		// 节点的最大路径和
		priceNewPath := node.Val + leftGain + rightGain

		// 更新答案
		maxSum = max(maxSum, priceNewPath)
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func maxPathSum1(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 节点的最大路径和
		priceNewPath := node.Val + leftGain + rightGain

		// 更新答案
		maxSum = max(maxSum, priceNewPath)

		// 返回节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func maxPathSum2(root *TreeNode) int {
	ans := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归左右
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 节点最大路径和
		priceNewPath := node.Val + leftGain + rightGain
		ans = max(ans, priceNewPath)
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return ans
}

func maxPathSum3(root *TreeNode) int {
	ans := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		//递归左右
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 节点最大路径和
		priceNewPath := node.Val + leftGain + rightGain
		ans = max(ans, priceNewPath)
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return ans
}
