package tow

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 199. 二叉树的右视图
func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			cur := queue[0]
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
