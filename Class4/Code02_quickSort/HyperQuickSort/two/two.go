package two

import (
	"math/rand"
	"time"
)

func sortArray(nums []int) []int {
	rand.Seed(time.Now().Unix())
	QuickSort(nums, 0, len(nums)-1)
	return nums
}

func QuickSort(nums []int, left, right int) {
	// 递归终止条件
	if left >= right {
		return
	}

	// 大于、小于边界
	lt, gt := left, right
	pivotIndex := rand.Intn(right-left+1) + left
	nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]
	pivot := nums[left]

	i := lt + 1
	for i <= gt {
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
	QuickSort(nums, left, lt-1)
	QuickSort(nums, gt+1, right)
}

func mergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	mid := n / 2
	left := nums[0:mid]
	right := nums[mid:]
	return merge(mergeSort(left), mergeSort(right))
}
func merge(left []int, right []int) []int {
	ans := []int{}
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			ans = append(ans, left[0])
			left = left[1:]
		} else {
			ans = append(ans, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		ans = append(ans, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		ans = append(ans, right[0])
		right = right[1:]
	}
	return ans
}
