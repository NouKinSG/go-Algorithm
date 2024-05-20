package One

// uniquePaths  机器人从左上角到右下角的路径数量
func unique(m int, n int) int {
	// 动态规划
	dp := make([][]int, m)

	// 初始化
	for i := range dp {
		dp[i] = make([]int, n)
		// 矩阵第一行第一列的值（路径数）都为 1
		dp[i][0] = 1
	}

	// 矩阵第一行第一列的值（路径数）都为 1
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// 从第二行到第一行，从第二列到第一列，依次计算路径数
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
