package main

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// 46.全排序
func permute(nums []int) [][]int {
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
		if !used[i] {
			used[i] = true
			cur = append(cur, nums[i])
			ans = back(nums, ans, cur, used)
			used[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	return ans
}
