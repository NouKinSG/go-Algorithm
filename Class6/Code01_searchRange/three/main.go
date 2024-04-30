package three

func searchRange(nums []int, target int) []int {
	//先找 start
	start := binary(nums, target)

	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := binary(nums, target+1) - 1
	return []int{start, end}
}

// 二分
func binary(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left++
		} else {
			right--
		}
	}
	return left
}
