package one

import "strconv"

// countAndSay 生成外观数列的第 n 个元素
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	// 定义一个生成x下一个数列的函数
	nextSequece := func(seq string) string {
		result := ""
		i := 0
		for i < len(seq) {
			count := 1
			for i+1 < len(seq) && seq[i] == seq[i+1] {
				count++
				i++
			}
			result += strconv.Itoa(count) + string(seq[i])
			i++
		}
		return result
	}

	// 定义初始 currentSay
	currentSay := "1"
	for i := 1; i < n; i++ {
		currentSay = nextSequece(currentSay)
	}
	return currentSay
}
