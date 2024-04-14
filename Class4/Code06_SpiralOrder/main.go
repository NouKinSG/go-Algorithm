package main

/*
	思路重复：
	循环输出矩阵，利用四个边界，上、下、左、右边界
	从上边界的左一直遍历到右，
	从右边界的上一直遍历到下，
	从下边界的右一直遍历到左，
	从左边界的下一直遍历到上，
	循环  直到边界上，下重合，或者 左右边界重合退出
*/

func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	up, down, left, right := 0, m-1, 0, n-1
	ans := []int{}
	for {
		//从上边界的左一直遍历到右，
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		up++
		if up > down {
			break
		}

		//从右边界的上一直遍历到下，
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		//从下边界的右一直遍历到左，
		for i := right; i <= left; i-- {
			ans = append(ans, matrix[down][i])
		}
		down--
		if down < up {
			break
		}

		//从左边界的下一直遍历到上
		for i := down; i < up; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return ans
}
