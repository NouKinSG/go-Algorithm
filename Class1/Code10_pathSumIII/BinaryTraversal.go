package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func PreorderTraversal(root *TreeNode) {
	// 前序遍历思路
	// 访问根节点 ——> 遍历左子树 ——> 遍历右子树
	if root == nil {
		return
	}

	// 访问 根节点
	fmt.Println(root.Val)

	// 遍历左子树
	PreorderTraversal(root.Left)

	// 遍历右子树
	PreorderTraversal(root.Right)
}

// 中序遍历   思路： 左子树 ——> 根节点 ——> 右子树
func InorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	//左
	InorderTraversal(root.Left)
	//根
	fmt.Println(root.Val)
	//右
	InorderTraversal(root.Right)
}

// 后序遍历
func PostorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	PostorderTraversal(root.Left)
	PostorderTraversal(root.Right)
	fmt.Println(root.Val)
}

// 层次遍历  思路：树的每一层，从左到右访问所有节点
func LevelorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	// 利用队列
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		cur := queue[0]   // 先获取将要出队元素
		queue = queue[1:] // 出队

		// 访问 队头
		fmt.Println(cur.Val)

		if cur.Left != nil {
			queue = append(queue, cur.Left) //左入队
		}

		if cur.Right != nil {
			queue = append(queue, cur.Right) // 右入队
		}
	}
}
