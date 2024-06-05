package one

// 1.两数之和
func twoSum(nums []int, target int) []int {
	ans := []int{}
	if len(nums) == 0 {
		return ans
	}

	indexMap := make(map[int]int)
	for i, v := range nums {
		if _, ok := indexMap[target-v]; ok {
			return []int{i, indexMap[target-v]}
		}
		indexMap[v] = i
	}
	return nil
}
