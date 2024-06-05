package four

func sortArray(nums []int) []int {
	// 递归终止条件
	if len(nums) <= 1 {
		return nums
	}
	//中间
	mid := len(nums) / 2
	// 左右
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])
	return merge(left, right)
}
func merge(left, right []int) []int {
	// 定义 答案 ans
	ans := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ans = append(ans, left[i])
			i++
		} else {
			ans = append(ans, right[j])
			j++
		}
	}
	// 把剩下的元素加入 ans
	ans = append(ans, left[i:]...)
	ans = append(ans, right[j:]...)
	return ans
}
