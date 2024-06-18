package Code03_evalRPN

import "strconv"

// 整体思路
// 1、题目给了数组，让所有元素依次入栈
// 2、入栈时，看看当前入栈元素是不是运算符
// 3、是就把前两个元素取出，然后做运算
// 4、运算的结果放入栈区
// 重复 1、2、3、4
// 最后栈区只有一个元素，就是答案

func evalRPN(tokens []string) int {
	//栈区
	var stack []int
	for _, token := range tokens {
		inttoken, err := strconv.Atoi(token)
		if err == nil {
			// 是数字
			stack = append(stack, inttoken)
		} else {
			// 是运算符.   err会不为nil

			// 弹出两个元素
			nums1, nums2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2] // 因为是数组模拟栈，所以这里是删除最后两个元素

			switch token {
			case "+":
				stack = append(stack, nums1+nums2)
			case "-":
				stack = append(stack, nums1-nums2)
			case "*":
				stack = append(stack, nums1*nums2)
			case "/":
				stack = append(stack, nums1/nums2)
			}
		}
	}
	return stack[0]
}
