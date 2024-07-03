package one

//func longestValidParentheses(s string) int {
//	ans := 0
//	dp := make([]int, len(s))
//
//	for i := 1; i < len(s); i++ {
//		if s[i] == ')' {
//			if s[i-1] == '(' {
//				if i >= 2 {
//					dp[i] = dp[i-2] + 2
//				} else {
//					dp[i] = 2
//				}
//			} else if i-dp[i-1] >0 && s[i-dp[i-1]-1] == '(' {
//				if i-dp[i-1] >= 2 {
//					dp[i] = dp[i-1]
//				}
//			}
//		}
//	}
//}

func longestValidParentheses(s string) int {
	stack := []int{}
	stack = append(stack, -1)
	maxLen := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			// 现在是 右括号，弹出栈顶元素，试图认为与左括号匹配
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 这里发现 栈元素未空
				// 所以 这个右括号没有能与它匹配的左括号
				stack = append(stack, i)
			} else {
				// 栈元素不为空，说明 有左括号与其匹配
				// 那我们需要更新最大长度
				length := i - stack[len(stack)-1]
				if length > maxLen {
					maxLen = length
				}
			}
		}
	}
	return maxLen
}

func longestValidParentheses1(s string) int {
	// 定义dp
	dp := make([]int, len(s))
	ans := 0
	// 遍历
	for i := 1; i < len(s); i++ {
		// 左括号不做考虑
		if s[i] == '(' {
			continue
		}

		// 右括号   dp[i-1] 表示  pre ~ i 之间，可能有的有效括号长度
		pre := i - 1 - dp[i-1]
		if pre >= 0 && s[pre] == '(' {
			dp[i] = 2 + dp[i-1]
			if pre > 0 {
				dp[i] += dp[pre-1]
			}
		}

		//ans = max(dp[i],ans)
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
