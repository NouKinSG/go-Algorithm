package two

/*
示例 1:

输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: nums = [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

*/

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// maxEnd 代表遍历到目前为止，最大的乘积
	// minEnd 代表遍历到目前为止，最小的乘积
	maxEnd, minEnd, maxAns := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxEnd, minEnd = minEnd, maxEnd
		}

		// 更新以当前元素结尾的最大乘积和最小乘积
		maxEnd = max(nums[i], maxEnd*nums[i])
		minEnd = min(nums[i], minEnd*nums[i])

		// 从上一次保存的Ans，与 当前最大乘积比较 取两者最大
		maxAns = max(maxAns, maxEnd)
	}
	return maxAns
}
