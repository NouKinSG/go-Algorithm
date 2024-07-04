package three

import "unicode"

func isPalindrome(s string) bool {
	// 定义过滤数组
	var filter []rune

	// 遍历字符
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			filter = append(filter, unicode.ToLower(char))
		}
	}

	// 左右指针遍历数组
	left, right := 0, len(filter)-1
	for left < right {
		if filter[left] != filter[right] {
			return false
		}
		left++
		right--
	}

	return true
}
