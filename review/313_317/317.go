package _13_317

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果有上和左
			if i-1 >= 0 && j-1 >= 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			} else {
				// 只有上
				if i-1 >= 0 {
					grid[i][j] += grid[i-1][j]
				}
				if j-1 >= 0 {
					grid[i][j] += grid[i][j-1]
				}
			}
		}
	}
	return grid[m-1][n-1]
}
