package main

// 14、最长公共前缀
func longestCommonPrefix(strs []string) string {
	m, n := len(strs), len(strs[0])
	if m == 0 {
		return ""
	}

	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
