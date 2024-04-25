package main

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	start := binarysearch(nums, target)
	// TODO  这里的判断不熟悉
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}

	// end 也忘记了
	end := binarysearch(nums, target+1) - 1
	return []int{start, end}
}

// 二分查找
func binarysearch(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if left < right {
			left = mid - 1
		} else {
			right = mid + 1
		}
	}
	return left
}
