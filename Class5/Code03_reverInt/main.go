package main

import "math"

func reverse(x int) (ans int) {

	// 不为 0，对 x 逐位遍历
	for x != 0 {
		// 如果 反转后整数(ans) 超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
		if ans <= math.MinInt32 || ans >= math.MaxInt32 {
			return 0
		}
		digital := x % 10
		x /= 10
		ans = ans*10 + digital
	}
	return
}
