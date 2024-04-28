package tow

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	// 判空
	if len(nums) == 0 {
		return []int{}
	}
	start := search(nums, target)
	if start == len(nums) || nums[start] != start {
		return []int{-1, 1}
	}
	end := search(nums, target)
	return []int{start, end}
}

// 二分查找
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if mid < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
