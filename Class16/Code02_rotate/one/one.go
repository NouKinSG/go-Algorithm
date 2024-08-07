package one

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n // 如果 k 大于数组长度，取模以避免多余的轮转

	count := 0 // 计数
	// 外层循环是 多个环的情况
	for i := 0; count < n; i++ {

		// 两个指针
		cur := i
		pre := i

		// 内层循环是 某个环的具体做。
		for {
			// 轮转后的下标
			next := (cur + k) % n
			// 交换
			nums[next], pre = pre, nums[next]
			cur = next
			// 换过一次计数
			count++
			if i == cur {
				break
			}
		}
	}
}
