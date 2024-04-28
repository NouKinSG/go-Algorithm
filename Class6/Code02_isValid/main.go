package code02isvalid

// 20. 有效的括号
func isValid(s string) bool {
	// 栈区

	stack := []rune{}

	// 映射闭合括号到开放括号
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// for _,char := range s{
	//     if openBracket,ok := bracketMap[char];ok{
	//         // 闭合口号，检查栈顶元素
	//         if len(stack) == 0 || stack[len(stack) - 1] != openBracket{
	//             return false
	//         }
	//         // 弹出栈顶元素
	//         stack = stack[:len(stack)-1]
	//     } else{
	//         // 如果是开放括号，放入栈中
	//         stack = append(stack,char)
	//     }
	// }
	// 如果栈为空，说明有效
	// return len(stack) == 0

	for _, char := range s {
		// 闭合括号，检查栈顶元素
		if openBracket, ok := bracketMap[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 如果开括号，放入栈中
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

// 20. 有效的括号
func isValid2(s string) bool {
	stack := []rune{}

	// 映射
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if openBracket, ok := bracketMap[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			// 弹出
			stack = stack[:len(stack)-1]
		} else {
			// 加入栈区
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
