package main

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))

	maxLen := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		maxLen = max(dp[i], maxLen)
	}
	return maxLen
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
