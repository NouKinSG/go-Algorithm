package two

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := math.MinInt32
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		LeftGain := max(maxGain(node.Left), 0)
		RightGain := max(maxGain(node.Right), 0)

		priceNewPath := node.Val + LeftGain + RightGain
		ans = max(ans, priceNewPath)
		return node.Val + max(LeftGain, RightGain)
	}
	maxGain(root)
	return ans
}
