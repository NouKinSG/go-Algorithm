package main

import "math"

/*
	描述：将一个整数反转，例如：x = 123，ans 输出 321，
	如果 反转后整数(ans) 超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。
	参数：传入一个整数
	返回值：返回

*/
func reverse(x int) (ans int) {

	// 不为 0，对 x 逐位遍历
	for x != 0 {
		// 如果 反转后整数(ans) 超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。
		if ans <= math.MinInt32/10 || ans >= math.MaxInt32/10 {
			return 0
		}
		digital := x % 10
		x /= 10
		ans = ans*10 + digital
	}
	return
}
