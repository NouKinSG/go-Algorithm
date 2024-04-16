package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// 定义答案
	ans := []int{}

	// 队列
	queue := []*TreeNode{}
	// root 入队
	queue = append(queue, root)
	for len(queue) != 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			cur := queue[0]
			// 出队
			queue = queue[1:]

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			if i == n-1 {
				ans = append(ans, cur.Val)
			}
		}
	}
	return ans
}
