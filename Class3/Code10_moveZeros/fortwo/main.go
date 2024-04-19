package main

// 283. 移动零

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
		}
		right++
	}
}
