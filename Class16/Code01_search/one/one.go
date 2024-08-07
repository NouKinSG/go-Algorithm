package one

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		// 找中间
		mid := (left + right) / 2
		if nums[mid] == target {
			// 题目是想要返回下标
			return mid
		}

		// 判断那部分有序
		if nums[left] <= nums[mid] {
			// 左半部分
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右半部分有序
			if nums[mid] < target && target < nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	// 没找到
	return -1
}
