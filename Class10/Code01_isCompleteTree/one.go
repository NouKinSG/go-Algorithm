package Code01_isCompleteTree

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	// 空树也算
	if root == nil {
		return true
	}

	//	使用切片队列
	queue := []*TreeNode{root}
	// 标识位
	readflag := false

	for len(queue) > 0 {
		//队列
		cur := queue[0]
		queue = queue[1:]
		if cur != nil {
			if readflag {
				return false
			}
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		} else {
			// 如果出现了空，那就以后必须为空
			readflag = true
		}
	}
	return true
}
