package four

import "sort"

func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	n := len(nums)
	ans := [][]int{}
	for i := 0; i < n; i++ {
		// 跳过重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 将寻找三个转变为寻找两个
		target := -nums[i]
		// 以前以后分别寻找
		left, right := i+1, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 跳过重复
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				// 值小了
				left++
			} else {
				right--
			}
		}
	}
	return ans
}
