package tow

// 37. 解数独
func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for c := byte('1'); c <= '9'; c++ {
					// 判断是否有效
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
				return false
			}
		}
	}
	return true
}
func isValid(board [][]byte, row, col int, c byte) bool {
	// 行有效
	for i := 0; i < 9; i++ {
		if board[row][i] == c {
			return false
		}
	}

	// 列有效
	for j := 0; j < 9; j++ {
		if board[j][col] == c {
			return false
		}
	}

	// 九宫格有效
	for i := 0; i < 9; i++ {
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == c {
			return false
		}
	}
	return true
}
