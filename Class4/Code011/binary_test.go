package main

import "testing"

// 中序遍历 dcade

// 还原成二叉树

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func binary(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return build(nums, 0, len(nums)-1)
}

func build(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	// 防止溢出
	mid := start + (end-start)/2
	root := &TreeNode{val: nums[mid]}
	root.left = build(nums, start, mid-1)
	root.right = build(nums, mid+1, end)
	return root
}

func TestBinry(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	root := binary(nums)

	// 验证 root 是否为 nil
	if root == nil {
		t.Errorf("Expected non-nil root, got nil")
	}

	// 进一步验证 root 的结构和值是否符合预期
	// ...

	// 其他断言和验证代码
	// ...
}
