package main

func longestConsecutive(nums []int) int {
	// 创建集合，去重
	numSet := make(map[int]bool)

	for _, val := range nums {
		numSet[val] = true
	}

	ans := 0
	for num := range numSet {
		// numSet[num - 1] 指的是 num的前面有没有数，判断num是不是起点，如果有会再集合中是true
		// 我们是在找起点，如果num前面有，说明不是起点，我们就直接跳过他
		// 找起点，如果num是起点，那么numSet[num-1]会是false，取反即可
		if !numSet[num-1] {
			current := num

			// 统计从起点开始，连续的有多少个
			currentArr := 1

			for numSet[current+1] {
				current++
				currentArr++
			}
			if ans < currentArr {
				ans = currentArr
			}
		}
	}
	return ans
}
