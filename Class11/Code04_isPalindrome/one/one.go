package one

import "unicode"

func isPalindrome(s string) bool {
	// 定义 过滤后的数组
	var filteredChars []rune

	// char 字符遍历并且
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {

			// 转成小写，加入过滤数组
			filteredChars = append(filteredChars, unicode.ToLower(char))
		}
	}

	// 双指针
	left, right := 0, len(filteredChars)-1
	for left < right {
		// 左边字符 不等于 右边字符      直接返回false
		if filteredChars[left] != filteredChars[right] {
			return false
		}
		// 移动指针
		left++
		right--
	}
	return true
}
