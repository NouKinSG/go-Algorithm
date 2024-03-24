package main

func isHappy(n int) bool {
	// 用map存储出现过的 n
	seen := make(map[int]bool)
	for n != 1 && !seen[n] {
		// 没出现过 map 中的，记录下来
		seen[n] = true
		n = getSum(n)
	}
	return n == 1
}

// 求 数位平方和
func getSum(n int) (ans int) {
	for n > 0 {
		digital := n % 10
		ans += digital * digital
		n /= 10
	}
	return ans
}
