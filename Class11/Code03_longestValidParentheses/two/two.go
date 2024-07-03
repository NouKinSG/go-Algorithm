package two

func longestValidParentheses(s string) int {
	// 判空
	if s == "" {
		return 0
	}
	ans := 0
	dp := make([]int, len(s))

	// 遍历
	for i := 1; i < len(s); i++ {
		// 左括号不做考虑
		if s[i] == '(' {
			continue
		}

		// 能来到这里的是 右括号
		pre := i - 1 - dp[i-1]
		if pre >= 0 && s[pre] == '(' {
			dp[i] = 2 + dp[i-1]
			if pre > 0 {
				dp[i] += dp[pre-1]
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
