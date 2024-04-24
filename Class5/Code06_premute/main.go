package main

func premute(nums []int) [][]int {
	var ans [][]int
	var cur []int
	used := make([]bool, len(nums))

	backtrack(nums, &ans, cur, used)

	return ans
}

func backtrack(nums []int, ans *[][]int, cur []int, used []bool) {
	if len(cur) == len(nums) {
		// 复制 cur，因为切片是引用类型,不能使用cur
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
			backtrack(nums, ans, cur, used)
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
}
