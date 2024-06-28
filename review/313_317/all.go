package _13_317

// 3.13接雨水
func trap(height []int) int {
	// 定义辅助变量
	n := len(height)
	ans := 0

	// 先判空
	if n == 0 {
		return ans
	}
	// 动态规划
	left, right := make([]int, n), make([]int, n)

	for i := 1; i < n; i++ {
		left[i] = max(height[i-1], left[i-1])
	}
	for i := n - 2; i > 0; i-- {
		right[i] = max(height[i+1], right[i+1])
	}

	for i := range height {
		short := min(left[i], right[i])
		if short > height[i] {
			ans += short - height[i]
		}
	}
	return ans
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
