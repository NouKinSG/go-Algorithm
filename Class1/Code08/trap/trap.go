package main

import (
	"fmt"
)

func trap2(height []int) int {
	n := len(height)

	// 定义 n长度，int类型切片
	left := make([]int, n)

	left[0] = height[0]

	// 定义int类型切片
	right := make([]int, n)
	right[0] = height[n-1]

	ans := 0

	// 第 i 棵柱子，左边最高柱子
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i-1])
	}

	fmt.Println("----------------")
	for i, v := range left {
		fmt.Printf("索引为：%v,值为：%v \n", i, v)
	}

	// 第 i 棵柱子，右边最高柱子
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i+1])
		short := min(left[i], right[i])

		//能接雨水
		if short > height[i] {
			ans += short - height[i]
		}
	}

	return ans
}

// max函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min函数
func min(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func trap(height []int) int {
	// 定义答案
	ans := 0
	n := len(height)
	// 保存 凹槽 两侧的最高柱子
	left := make([]int, n)
	right := make([]int, n)

	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i-1])
	}

	for i := n - 2; i > 0; i-- {
		right[i] = max(right[i+1], height[i+1])
	}

	for i := 0; i < n; i++ {
		short := min(left[i], right[i])
		if short > height[i] {
			ans += short - height[i]
		}
	}
	return ans
}

func trap3(height []int) int {
	ans := 0
	n := len(height)
	left := make([]int, n)
	right := make([]int, n)

	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i-1])
	}

	for i := n - 2; i > 0; i-- {
		right[i] = max(right[i+1], height[i+1])
	}

	for i := 1; i < n; i++ {
		short := min(left[i], right[i])
		if short > height[i] {
			ans += short - height[i]
		}
	}

	return ans
}

// 尝试最优解
func trap4(height []int) (ans int) {
	ans = 0
	n := len(height)
	if n == 0 {
		return ans
	}
	arrLeft := make([]int, n)
	arrRight := make([]int, n)
	for i, j := 1, n-2; i < n; i++ {
		arrLeft[i] = max(arrLeft[i-1], height[i-1])
		arrRight[j] = max(arrRight[j+1], height[j+1])
		j--
	}

	for index, val := range height {
		short := min(arrLeft[index], arrRight[index])
		if short > val {
			ans += short - val
		}
	}

	return ans
}
