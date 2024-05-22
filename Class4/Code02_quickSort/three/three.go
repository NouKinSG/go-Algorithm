package three

import "math/rand"

func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	// 递归终止条件
	if left >= right {
		return
	}

	lt, gt := left, right
	pivotIndex := rand.Intn(right-left+1) + left
	nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]
	pivot := nums[left]

	i := lt + 1
	for i <= right {
		if nums[i] < pivot {
			nums[i], nums[lt] = nums[lt], nums[i]
			lt++
			i++
		} else if nums[i] > right {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		} else {
			i++
		}
	}
	quickSort(nums, left, lt-1)
	quickSort(nums, gt+1, right)
}
