package one

/*
解题思路：
	动态规划：动态二维数组dp
dp[i][j] 表示 s[i:j+1] 是否是回文。
如果 dp[i][j] 为true，那么 s[i:j+1] 是回文
状态转移方程： dp[i][j] =dp[i + 1][j - 1] && (s[i] == s[j])

先处理单个字符
再处理两个字符
最后处理两个以上字符

*/

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	// 定义dp
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	start := 0
	maxLength := 1
	// 单个字符是回文
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}

	// 处理两个字符
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLength = 2
		}
	}

	// 处理两个以上
	for length := 3; length <= len(s); length++ {
		// 剩余
		for i := 0; i <= len(s)-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start = i
				maxLength = length
			}
		}
	}
	return s[start : start+maxLength]
}
