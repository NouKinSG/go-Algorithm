package three

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(0, maxGain(node.Left))
		right := max(0, maxGain(node.Right))

		prices := node.Val + left + right
		ans = max(ans, prices)
		return node.Val + max(left, right)
	}
	maxGain(root)
	return ans
}
