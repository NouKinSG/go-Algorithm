package two

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
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
				// 偶数 正序
				temp = append(temp, node.Val)
			} else {
				temp = append([]int{node.Val}, temp...)
			}
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}

			queue = nextQueue
			ans = append(ans, temp)
			level++
		}
	}
	return ans
}
