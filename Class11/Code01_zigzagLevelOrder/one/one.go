package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 103. 二叉树的锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		nextQueue := []*TreeNode{}
		temp := []int{}

		for _, node := range queue {
			if level%2 == 0 {
				// 偶数层，正序
				temp = append(temp, node.Val)
			} else {
				// 奇数层，逆序
				temp = append([]int{node.Val}, temp...)
			}
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}

		queue = nextQueue
		ans = append(ans, temp)
		level++
	}
	return ans
}
