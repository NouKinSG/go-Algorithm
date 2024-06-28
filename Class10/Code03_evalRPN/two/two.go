package two

import "strconv"

func evalRPN(tokens []string) int {
	// 栈区
	var stack []int
	for _, token := range tokens {
		intToken, err := strconv.Atoi(token)
		if err == nil {
			// 是数字
			// 入栈
			stack = append(stack, intToken)
		} else {
			// 不是数字，是 符号

			// 取出前两个数字
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
