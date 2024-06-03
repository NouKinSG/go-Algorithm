package Code01_heapSort

// 堆排序函数
func sortArray(nums []int) []int {
	n := len(nums)
	//构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	// 一个一个从堆中取出元素
	for i := n - 1; i >= 0; i-- {
		// 当前堆的根（最大元素）放到数组末尾
		nums[0], nums[i] = nums[i], nums[0]

		//调整堆
		heapify(nums, i, 0)
	}
	return nums
}

func heapify(nums []int, n, i int) {
	largest := i     // 初始化最大值为根节点
	left := 2*i + 1  // 左子节点
	right := 2*i + 2 // 右子节点
	// 如果左子节点大于根节点
	if left < n && nums[left] > nums[largest] {
		largest = left
	}
	// 如果右子节点大于最大值
	if right < n && nums[right] > nums[largest] {
		largest = right
	}

	// 如果最大值不是根节点
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i] // 交换

		// 递归调整堆栈段
		heapify(nums, n, largest)
	}
}
