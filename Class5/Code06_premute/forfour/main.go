package forfour

func permute(nums []int) [][]int {
	var ans [][]int
	var cur []int
	used := make([]bool, len(nums))
	ans = back(nums, ans, cur, used)
	return ans
}

func back(nums []int, ans [][]int, cur []int, used []bool) [][]int {
	if len(nums) == len(cur) {
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
