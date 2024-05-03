package main

// exit 函数检查在二维网格中是否可以找到给定的单词
func exist(board [][]byte, word string) bool {
	// 遍历二维网格的每一个元素
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// 从每一个元素开始，尝试使用 dfs 函数查找单词
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

// dfs 函数执行深度优先搜索算法来匹配单词
func dfs(board [][]byte, i, j int, word string, k int) bool {
	// 如果当前字符索引 k 等于单词长度，表示找到了完整单词
	if k == len(word) {
		return true
	}
	// 检查当前坐标是否越界，或者当前字符不匹配
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[k] {
		return false
	}

	// 临时保存当前字符，并将当前位置标记为已访问（使用'#')
	temp := board[i][j]
	board[i][j] = '#'
	// 递归遍历四个可能性方向
	found := dfs(board, i+1, j, word, k+1) ||
		dfs(board, i-1, j, word, k+1) ||
		dfs(board, i, j+1, word, k+1) ||
		dfs(board, i, j-1, word, k+1)
	// 还原当前位置的字符
	board[i][j] = temp
	return found
}
