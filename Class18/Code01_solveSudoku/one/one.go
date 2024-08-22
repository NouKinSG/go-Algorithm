package one

func solveSudoku(board [][]byte) {
	// 自己玩数独游戏解数独步骤
	// 1、从 1 ~ 9 依次看每一个数字的分布。
	// 2、某一个数字的时候，行、列、九宫格就不能有这个数字。

	// 翻译到程序可以这样，我先把全部的空格设置成 待填空可选数字 1 ~ 9
	// 2、然后依次遍历数字每遇到一个数字，就把这个数字在的 行、列、 九宫格中的空格去除这个数字。
	// 这样遍历完了之后，答案不会全部出来，但是会出现某些空格只剩下一个数字可以填
	// 填完只有一个数字的全部空格后，重复 2、

	solve(board)

}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				// 这是空格，尝试填入 1 ~ 9
				for c := byte('1'); c <= '9'; c++ {
					// 检查数字 c 是否可以放在当前格子
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
				// 如果 '1' 到 '9' 都不能放在当前格子，返回 false 进行回溯
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, c byte) bool {
	// 列存在，无效
	for i := 0; i < 9; i++ {
		if board[i][col] == c {
			return false
		}
	}

	// 行存在，无效
	for j := 0; j < 9; j++ {
		if board[row][j] == c {
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
