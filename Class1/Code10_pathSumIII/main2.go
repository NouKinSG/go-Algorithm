package main

func pathSum2(root *TreeNode, targetSum int) int {
	ans := 0
	if root == nil {
		return ans
	}

	// 以 root 为起点，看看和为targetSum的路径数
	ans += pathFromNode2(root, targetSum)
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)
	return ans
}

// 以 给定节点为起点，和为targetSum的路径数。
func pathFromNode2(node *TreeNode, targetSum int) int {
	count := 0
	if node == nil {
		return count
	}

	if node.Val == targetSum {
		count++
	}

	// 继续往子树看看，还有没有，但是子树找的目标值 = targetSum - node.val
	count += pathFromNode2(node.Left, targetSum-node.Val)
	count += pathFromNode2(node.Right, targetSum-node.Val)
	return count
}
