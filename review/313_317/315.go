package _13_317

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 437. 路径总和 III
func pathSum(root *TreeNode, targetSum int) int {
	ans := 0
	if root == nil {
		return ans
	}
	ans += pathSumForNode(root, targetSum)
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)
	return ans
}

func pathSumForNode(node *TreeNode, sum int) int {
	count := 0
	if node == nil {
		return count
	}
	if node.Val == sum {
		count++
	}
	count += pathSumForNode(node.Left, sum-node.Val)
	count += pathSumForNode(node.Right, sum-node.Val)
	return count
}
