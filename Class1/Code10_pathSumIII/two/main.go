package two

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
	ans += pathSumFromNode(root, targetSum)
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)
	return ans
}

func pathSumFromNode(node *TreeNode, targetSum int) int {
	count := 0
	if node == nil {
		return count
	}

	if node.Val == targetSum {
		count++
	}
	count += pathSumFromNode(node.Left, targetSum-node.Val)
	count += pathSumFromNode(node.Right, targetSum-node.Val)
	return count
}
