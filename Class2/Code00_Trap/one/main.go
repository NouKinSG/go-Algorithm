package main

func trap(height []int) int {
	ans := 0

	n := len(height)
	if n == 0 {
		return ans
	}
	left, right := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		left[i] = max(height[i-1], left[i-1])
	}

	for i := n - 2; i > 0; i-- {
		right[i] = max(height[i+1], right[i+1])
	}

	for i := range height {
		short := min(left[i], right[i])
		if short > height[i] {
			ans += short - height[i]
		}
	}

	return ans
}
