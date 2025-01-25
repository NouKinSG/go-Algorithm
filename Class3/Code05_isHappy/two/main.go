package two

//202. 快乐数

//编写一个算法来判断一个数 n 是不是快乐数。
//
//「快乐数」 定义为：
//
//对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
//然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
//如果这个过程 结果为 1，那么这个数就是快乐数。
//如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

func isHappy(n int) bool {

	// 使用 map 存储出现过的 n
	seen := make(map[int]bool)
	for n != 1 && !seen[n] {

		// 先记录没有出现过的 n
		seen[n] = true
		n = getSum(n)
	}

	return n == 1
}

func getSum(n int) int {
	res := 0
	for n > 0 {
		digits := n % 10
		res += digits * digits
		n /= 10
	}
	return res
}
