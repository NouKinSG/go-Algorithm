package one

func sortColors(nums []int) {
	if nums == nil {
		return
	}

	// 准备几个指针
	var (
		cur = 0
		l   = 0
		r   = len(nums) - 1
	)

	for cur <= r {
		if nums[cur] < 1 {
			// 小于 1，
			nums[cur], nums[l] = nums[l], nums[cur]
			cur++
			l++
		} else if nums[cur] == 1 {
			cur++
		} else {
			nums[cur], nums[r] = nums[r], nums[cur]
			r--
		}
	}
}
