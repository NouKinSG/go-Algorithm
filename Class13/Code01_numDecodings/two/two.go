package two

// 91. 解码方法
func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s == "" {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1 // 空字符串
	dp[1] = 1 // 1个字符串

	for i := 2; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
