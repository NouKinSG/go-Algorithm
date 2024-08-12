package one

func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '1' // '1' ~ '9' 字符，有没有出现过，我们用 数组的索引表示，但是索引只有 0 ~ 8 所以减 1
			boxIndex := (i/3)*3 + j/3

			// 只要有任何一个 存在了，就说明这个数独没有效
			if rows[i][num] || cols[j][num] || boxes[boxIndex][num] {
				return false
			}

			rows[i][num] = true
			cols[j][num] = true
			boxes[boxIndex][num] = true
		}
	}
	return true
}
