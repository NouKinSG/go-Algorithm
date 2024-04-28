package main

// 20. 有效的括号
func isValid1(s string) bool {
	// 栈区
	stack := []rune{}

	// 映射闭合括号到开放括号
	bracketMap := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, char := range s {

		// 如果char是闭合括号，bracketMap[char] 是开放括号
		if openBracket, ok := bracketMap[char]; ok {

			if len(stack) == 0 || openBracket != stack[len(stack)-1] {
				// 如果栈区为 0，或者 有闭合括号，但是没开放括号
				// 直接 return
				return false
			}
			// 判断完之后弹出
			stack = stack[:len(stack)-1]
		} else {
			// 入栈
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
