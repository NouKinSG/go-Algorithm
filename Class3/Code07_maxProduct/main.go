package main

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxEnd, minEnd, maxResult := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		// 如果当前元素是负数，交换最大最小值
		if nums[i] < 0 {
			maxEnd, minEnd = minEnd, maxEnd
		}
		// 更新以元素结尾的最大乘积和最小乘积
		maxEnd = max(nums[i], maxEnd*nums[i])
		minEnd = min(nums[i], minEnd*nums[i])

		//更新最大乘积
		maxResult = max(maxResult, maxEnd)
	}
	return maxResult
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
