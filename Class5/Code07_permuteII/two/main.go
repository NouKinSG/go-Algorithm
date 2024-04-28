package main

import "sort"

// 47.全排序 II
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	var cur []int
	used := make([]bool, len(cur))

	ans = back(nums, ans, cur, used)
	return ans
}

func back(nums []int, ans [][]int, cur []int, used []bool) [][]int {
	if len(cur) == len(nums) {
		temp := make([]int, len(cur))
		copy(temp, cur)
		ans = append(ans, temp)
	}

	for i := range nums {
		// 前置校验
		// 当前的 数字使用过， 或者前一个数字未使用过并且当前的数字和上一个数字相同。直接跳过
		if used[i] || i > 0 && !used[i-1] && nums[i] == nums[i-1] {
			continue
		}
		if !used[i] {
			cur = append(cur, nums[i])
			used[i] = true
			ans = back(nums, ans, cur, used)
			used[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	return ans
}
