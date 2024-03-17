package main

func minPathSum(grid [][]int) int {
	ans := 0
	if len(grid) == 0 {
		return ans
	}

	row, col := len(grid), len(grid[0])
	// gridSun := make([][]int,row)
	// copy(gridSun,grid)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// 两条路径
			if i-1 >= 0 && j-1 >= 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			} else {
				// 只有左边一条
				if i-1 >= 0 {
					grid[i][j] += grid[i-1][j]
				}
				// 只有右边一条
				if j-1 >= 0 {
					grid[i][j] += grid[i][j-1]
				}
			}
			ans = grid[i][j]
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
