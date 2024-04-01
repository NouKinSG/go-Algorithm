package main

import (
	"math/rand"
	"time"
)

func sortArray(nums []int) []int {
	// 随机种子
	rand.Seed(time.Now().Unix())
	QuickSort(nums, 0, len(nums)-1)
	return nums
}

func QuickSort(nums []int, left int, right int) {
	// 递归终止条件
	if left >= right {
		return
	}

	// 小于基准值和大于基准值的边界
	lt, gt := left, right
	// 随机选择一个基准值，并将其与左边界元素交换
	pivotIndex := rand.Intn(right-left+1) + left
	nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]
	pivot := nums[left]

	// 从左边界的下一个元素开始遍历数组
	i := left + 1
	for i <= gt {
		// 如果当前元素小于基准值，将其与 lt 位置的元素交换，lt 和 i 都向右移动
		if nums[i] < pivot {
			nums[i], nums[lt] = nums[lt], nums[i]
			lt++
			i++
		} else if nums[i] > pivot {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		} else {
			i++
		}
	}
	QuickSort(nums, left, lt-1)
	QuickSort(nums, gt+1, right)
}
