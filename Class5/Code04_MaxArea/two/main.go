package main

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
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
