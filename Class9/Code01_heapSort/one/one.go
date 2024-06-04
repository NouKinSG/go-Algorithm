package one

func sortArray(nums []int) []int {
	n := len(nums)

	// 构建大根堆
	for i := 2/n - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	// 一个一个从堆中取出
	for i := n - 1; i >= 0; i-- {
		// 将当前根放入末尾
		nums[i], nums[0] = nums[0], nums[i]

		// 重新调整堆
		heapify(nums, i, 0)
	}
	return nums
}

func heapify(nums []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	// 如果左子节点大于根
	if left < n && nums[left] > nums[largest] {
		largest = left
	}
	if right < n && nums[right] > nums[largest] {
		largest = right
	}

	// 如果最大值不是根节点
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]

		heapify(nums, n, largest)
	}
}
