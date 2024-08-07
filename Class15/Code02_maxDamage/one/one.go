package one

// maxDamage 函数计算最多能造成给敌人的伤害
func maxDamage(maxCost int, costs []int, damages []int) int {
	// 思路：动态规划
	n := len(costs)
	// dp[i][j] 表示前i个物品，花费不超过j的情况下，能造成的最大伤害
	dp := make([][]int, n+1)
	for i := range dp {
		// dp[i] 表示 前i个物品
		dp[i] = make([]int, maxCost+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= maxCost; j++ {
			// 不选第i个物品，伤害不变
			dp[i][j] = dp[i-1][j]
			if j >= costs[i-1] {
				// 选第i个物品，伤害增加
				dp[i][j] = max(dp[i][j], dp[i-1][j-costs[i-1]]+damages[i-1])
			}
		}
	}
	return dp[n][maxCost]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
