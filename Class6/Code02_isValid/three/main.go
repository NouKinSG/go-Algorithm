package main

func isValid(s string) bool {
	// 栈区
	stack := []rune{}
	// 哈希映射
	bracketMap := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, char := range s {
		if openBracket, ok := bracketMap[char]; ok {
			// 有闭合符号，找开放符号
			if len(stack) == 0 || openBracket != stack[len(stack)-1] {
				return false
			}
			// 弹出
			stack = stack[:len(stack)-1]
		} else {
			// 入栈
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
