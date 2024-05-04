package tow

/*
79. 单词搜索

给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。
如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母不允许被重复使用。

*/

// 79. 单词搜索
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j int, word string, k int) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[k] {
		return false
	}

	temp := board[i][j]
	board[i][j] = '#'
	found := dfs(board, i-1, j, word, k+1) ||
		dfs(board, i+1, j, word, k+1) ||
		dfs(board, i, j+1, word, k+1) ||
		dfs(board, i, j-1, word, k+1)
	board[i][j] = temp
	return found
}
