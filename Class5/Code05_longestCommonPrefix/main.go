package main

// 14. 最长公共前缀

// 寻找公共前缀  比如 strs = ["flower","flow","flight"]
// 1. 用一个 字符的第 i 个字符与其他字符的第i个检查。
// 2. 比如用  flower，先遍历 flower ，这时 i 表示 flower第i个字
// 3. 然后 第 i 个字符与，其他字符检查。
// 如果 i = flowe 最后一个字符时，返回，或者 第 i 个 字符与其他字段不等时，返回

// 14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	m, n := len(strs), len(strs[0])

	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
