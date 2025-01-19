package main

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// 有上边界和左边界
			if i-1 >= 0 && j-1 >= 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			} else {
				// 左
				if j-1 >= 0 {
					grid[i][j] += grid[i][j-1]
				}
				if i-1 >= 0 {
					grid[i][j] += grid[i-1][j]
				}
			}
		}
	}
	return grid[row-1][col-1]
}
