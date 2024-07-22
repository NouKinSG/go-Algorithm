package one

import "sort"

func groupAnagrams(strs []string) [][]string {
	// 定义map
	m := map[string][]string{}
	for _, str := range strs {
		// 转成byte数组
		b := []byte(str)
		// 排序
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		// 排序后的字符串数组
		sortedStr := string(b)
		// 存入map
		m[sortedStr] = append(m[sortedStr], str)
	}

	// 构建结果
	ans := make([][]string, 0, len(m))
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
