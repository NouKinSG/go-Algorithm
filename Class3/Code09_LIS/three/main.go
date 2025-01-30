package three

func maxLengthLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	maxLength := 1

	for i := 1; i < len(nums); i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], dp[j]+1)
		}
		maxLength = max(dp[i], maxLength)
	}
	return maxLength
}
