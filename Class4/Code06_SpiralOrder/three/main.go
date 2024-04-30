package three

// 54. 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return []int{}
	}
	ans := []int{}
	up, down, left, right := 0, m-1, 0, n-1

	for {
		// 上
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		// 右
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		// 下
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
		down--
		if down < up {
			break
		}

		// 左
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return ans
}
