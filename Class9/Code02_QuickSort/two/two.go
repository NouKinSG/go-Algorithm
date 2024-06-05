package two

import (
	"math/rand"
	"time"
)

func sortArray(nums []int) []int {
	// 随机种子
	rand.Seed(time.Now().Unix())
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	// 递归终止条件
	if left <= right {
		return
	}

	lt, gt := left, right
	pivotIndex := rand.Intn(right-left+1) + left
	nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]
	pivot := nums[left]

	// 定义i
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

	// 递归
	quickSort(nums, left, lt-1)
	quickSort(nums, gt+1, right)
}
