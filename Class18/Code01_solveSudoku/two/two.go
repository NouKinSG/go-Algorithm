package two

func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				// 空格，填入 1 ~ 9
				for c := byte('1'); c < '9'; c++ {
					// 检查 c 能否放入格子
					if isValid(board, i, j, c) {
						board[i][j] = c

						// 递归
						if solve(board) {
							return true
						} else {
							// 刚才的数字不对
							// 回溯
							board[i][j] = '.'
						}
					}
				}
				// 1 ~ 9 都不能放，放回false
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == c {
			return false
		}
	}

	for j := 0; j < 9; j++ {
		if board[j][col] == c {
			return false
		}
	}

	// 当前九宫格存在了，无效
	for i := 0; i < 9; i++ {
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == c {
			return false
		}
	}
	return true
}
