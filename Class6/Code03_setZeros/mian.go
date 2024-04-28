package main

func setZeroes(matrix [][]int) {
	// 判空
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return
	}

	// 标志位
	rowFlag := false
	colFlag := false

	// 遍历第一行和第一列，如果有0，对应符号位值为 true
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			rowFlag = true
			break
		}
	}
	for j := 0; j < n; j++ {
		if matrix[j][0] == 0 {
			colFlag = false
			break
		}
	}

	// 遍历数组  将其对应的标志位值为 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	// 根据标志位，将矩阵元素值为 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 最后处理 第一行和第一列
	if rowFlag {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	for colFlag {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
}
