package three

// 73.矩阵置零
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return
	}

	rowFlag := false
	colFlag := false

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			rowFlag = true
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			colFlag = true
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if rowFlag {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	if colFlag {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
}
