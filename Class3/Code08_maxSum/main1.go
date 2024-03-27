package main

func mainSum1(nums []int) int {
	ans := 0
	if len(nums) == 0 {
		return ans
	}

	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if ans < nums[i] {
			ans = nums[i]
		}
	}
	return ans
}
