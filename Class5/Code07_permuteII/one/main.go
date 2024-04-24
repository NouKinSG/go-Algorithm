package main

import "sort"

// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

// 47.全排序
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	var cur []int
	used := make([]bool, len(nums))
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
