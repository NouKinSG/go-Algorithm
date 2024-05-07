package sortArrayToBST

// 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedArrayToBST 接收一个有序数组，并返回一个高度平衡的二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	ans := &TreeNode{}
	if len(nums) == 0 {
		return ans
	}
	return buildBST(nums, 0, len(nums)-1)
}

// buildBST 递归函数，用于构建二叉搜索树
func buildBST(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2 // 防止溢出
	node := &TreeNode{Val: nums[mid]}
	node.Left = buildBST(nums, left, mid-1)
	node.Right = buildBST(nums, mid+1, right)
	return node
}
