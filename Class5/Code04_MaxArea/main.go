package main

// 中等
// 提示
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。

// 思路重复

//11. 盛最多水的容器
func maxArea(height []int) int {
	ans := 0
	if len(height) == 0 {
		return ans
	}

	left, right := 0, len(height)-1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		ans = max(ans, area)
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
