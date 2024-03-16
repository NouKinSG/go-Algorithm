package main

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 当前节点作为起点的路径数
	rootCount := pathsFromNode(root, targetSum)

	// 左子树的路径和
	leftCount := pathSum(root.Left, targetSum)

	// 右子树的路径和
	rightCount := pathSum(root.Right, targetSum)
	n := rootCount + leftCount + rightCount
	return n
}

// 从给定节点开始，向下寻找总和为targetSum的路径数
func pathsFromNode(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}

	// 当前的节点值是否等于targetSum，等于为 1，否则为 0
	count := 0
	if node.Val == sum {
		count = 1
	}
	// 继续向下找
	count += pathsFromNode(node.Left, sum-node.Val)
	count += pathsFromNode(node.Right, sum-node.Val)
	return count
}

func pathSumIII(root *TreeNode, target int) int {
	ans := 0
	if root == nil {
		return ans
	}

	ans += pathFromNode(root, target)
	ans += pathSumIII(root.Left, target)
	ans += pathSumIII(root.Right, target)

	return ans
}

// 从给定节点开始，向下寻找 和 为 target 的路径数
func pathFromNode(node *TreeNode, Sum int) int {
	count := 0

	if node == nil {
		return 0
	}

	if node.Val == Sum {
		count = 1
	}

	count += pathFromNode(node.Left, Sum-node.Val)
	count += pathFromNode(node.Right, Sum-node.Val)

	return count
}

func main() {
	// Your code here
	a := make(chan int)
}
