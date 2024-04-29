package main

/* 思路重复

题目：给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
1、先判空
2、用标志位标记 记住第一行和第一列有没有为0的
3、遍历一遍矩阵，如果要是第 i 行的第 j 列的元素为0，
	就把 第 0 行的第 j列元素置为0，第 i 行的 第 0 个元素置为 0
4、在遍历一遍举证，根据第一行和第一列，为0的元素，把同行，同列全置为 0
5、最后再根据 标志位把置为 0 行和列，元素都置为 0。

*/

// 73.矩阵置零
func setZeroes(matrix [][]int) {
	// 第一步、判空
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return
	}

	// 第二部、标记 首行首列
	rowFlag := false
	colFlag := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			rowFlag = true
			break
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			colFlag = true
			break
		}
	}

	// 第三步、遍历矩阵,对应首行首列置0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 第四步，一行一列置0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 第五步、根据一开始的 标志位，把 首行首列置0
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
