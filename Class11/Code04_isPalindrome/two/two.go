package two

import "unicode"

func isPalindrome(s string) bool {
	// 定义一个过滤后的切片
	var filter []rune

	// char 字符遍历
	for _, char := range s {
		// 遍历之后看字符 保留 字母和数字字符
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			// 转成小写 加入过滤切片
			filter = append(filter, unicode.ToLower(char))
		}
	}

	// 使用头尾双指针 遍历
	left, right := 0, len(filter)-1
	for left < right {
		// 判值
		if filter[left] != filter[right] {
			return false
		}
		// 移动指针
		left++
		right--
	}

	return true
}
