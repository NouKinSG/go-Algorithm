package two

func isValidSudoku(board [][]byte) bool {

	var rows, cols, boxes [9][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '1'
			boxIndex := (i/3)*3 + j/3

			if rows[i][num] || cols[j][num] || boxes[boxIndex][num] {
				return false
			}

			rows[i][num] = true
			rows[j][num] = true
			boxes[boxIndex][num] = true
		}
	}
	return true
}
