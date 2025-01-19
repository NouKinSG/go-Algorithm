package two

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

	for j := n - 2; j > 0; j-- {
		right[j] = max(height[j+1], right[j+1])
	}

	for j := range height {
		short := min(left[j], right[j])
		if short > height[j] {
			ans += short - height[j]
		}
	}
	return ans
}
