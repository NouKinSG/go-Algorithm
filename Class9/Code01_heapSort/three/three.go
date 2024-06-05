package three

func sortArray(nums []int) []int {
	n := len(nums)
	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	// 一个一个取出
	for i := n - 1; i >= 0; i-- {
		// 当前堆的根放入末尾
		nums[0], nums[i] = nums[i], nums[0]

		// 调整堆
		heapify(nums, i, 0)
	}
	return nums
}

func heapify(nums []int, n, i int) {
	largest := i
	left := 2*i + 1 // 左子节点
	right := 2*i + 2
	// 如果子节点大
	if left < n && nums[left] > nums[largest] {
		largest = left
	}
	if right < n && nums[right] > nums[largest] {
		largest = right
	}

	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]

		// 递归调整堆栈段
		heapify(nums, n, largest)
	}
}
