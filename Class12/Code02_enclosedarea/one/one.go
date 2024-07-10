package one

// 我的想法是，从题目出发。
// 题目给了，连接，区域，围绕，
// 怎么用代表表达出这三个操作。
// 然后我想，我先连接那些 'O'，形成区域，
// 最后我再看这些O区域中有没有在边缘的，如果区域中没有任何单元格位于 board 边缘，则该区域被 'X' 单元格围绕。
func solve(board [][]byte) {

}

// 连接
func connectBoundaryO(board [][]byte) {
	m, n := len(board), len(board[0])
	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
			return
		}
		// 可以连接
		board[x][y] = 'A'
		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}

	// 遍历边界
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(i, n-1)
		}
	}

	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(m-1, j)
		}
	}

}

// 区域

// 围绕
