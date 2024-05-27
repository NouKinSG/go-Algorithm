package four

import (
	"math/rand"
	"time"
)

func sortArray(nums []int) []int {
	// 随机数种子
	rand.Seed(time.Now().Unix())
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// 快速排序
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	lt, gt := left, right
	pivotIndex := rand.Intn(right-left+1) + left
	nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]
	pivot := nums[left]

	i := lt + 1
	for i <= gt {
		if nums[i] < pivot {
			nums[i], nums[lt] = nums[lt], nums[i]
			i++
			lt++
		} else if nums[i] > pivot {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		} else {
			i++
		}
	}
	quickSort(nums, left, lt-1)
	quickSort(nums, gt+1, right)
}
