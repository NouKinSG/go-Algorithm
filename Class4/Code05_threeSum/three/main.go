package main

import "sort"

// 15.三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := [][]int{}
	for i := 0; i < n; i++ {
		// 跳过重复元素
		for n > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left, right := i+1, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 跳过重复的 元素
				for left < right && nums[i] == nums[i+1] {
					left++
				}
				for left < right && nums[i] == nums[i-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				// 值大了，right左移动
				right--
			}
		}
	}
	return ans
}
