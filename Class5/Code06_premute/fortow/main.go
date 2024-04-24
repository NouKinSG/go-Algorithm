package main

// 46.全排序   回溯法

// 思路：
/*
	// 题目给了 n 个数，需要我们从左至右一次填入给定的 n个数，并且每个数只能使用一次

	1、定义一个 backtarck(nums []int, results *[][]int, current []int, used []bool) 函数，
	used 用来标记使用过的数， cur 用来保存已经使用数的序列，results是答案

	2、如果 len(cur) == len(nums) 说明数用完了，可以将 cur 装入答案results中，递归结束
	3、如果 len(cur) < len(nums)，说明还有数没被用完，检查所有数，找到未被使用的，填入 cur,并标记使用

	4、回溯后，将上一次使用的数，标记为 未使用过。

*/

// permute 生成数组 nums 的所有可能全排列
func permute(nums []int) [][]int {
	var ans [][]int
	var cur []int
	used := make([]bool, len(nums))
	return backtrack(nums, ans, cur, used)
}

func backtrack(nums []int, ans [][]int, cur []int, used []bool) [][]int {
	if len(cur) == len(nums) {
		temp := make([]int, len(cur))
		copy(temp, cur)
		ans = append(ans, temp)
		return ans
	}

	for i := range nums {
		if !used[i] {
			used[i] = true
			cur = append(cur, nums[i])
			ans = backtrack(nums, ans, cur, used)
			used[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	return ans
}

func backtrack1(nums []int, ans *[][]int, cur []int, used []bool) {
	if len(cur) == len(nums) {
		// 复制 cur，因为切片是引用类型
		temp := make([]int, len(cur))
		copy(temp, cur)
		*ans = append(*ans, temp)
		return
	}

	for i, num := range nums {
		if !used[i] {
			// 标记当前数字为已使用
			used[i] = true
			// 加入当前数字到 cur
			cur = append(cur, num)
			// 继续递归下一个数字
			backtrack1(nums, ans, cur, used)
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
}
