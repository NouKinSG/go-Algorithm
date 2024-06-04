package two

func sortArray(nums []int) []int {
	n := len(nums)
	// 构建根堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	// 一个一个加入根堆
	for i := n - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]

		heapify(nums, i, 0)
	}
	return nums
}

func heapify(nums []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// 判断左右节点
	if left < n && nums[left] > nums[largest] {
		largest = left
	}
	if right < n && nums[right] > nums[largest] {
		largest = right
	}

	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]

		heapify(nums, n, largest)
	}
}
