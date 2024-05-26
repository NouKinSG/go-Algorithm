package two

import "strconv"

// countAndSay 生成外观数列的第 n 个元素
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	// 定义一个生成 x 下一个数列的函数
	nextSequece := func(s string) string {
		result := ""
		i := 0
		if i < len(s) {
			count := 1
			for i+1 < len(s) && s[i] == s[i+1] {
				count++
				i++
			}
			result += strconv.Itoa(count) + string(s[i])
			i++
		}
		return result
	}

	// 定义一个初始画的 字符
	currentSay := "1"
	for i := 1; i < n; i++ {
		currentSay = nextSequece(currentSay)
	}
	return currentSay
}
