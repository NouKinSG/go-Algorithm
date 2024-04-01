package main

func sortArray(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)
	return nums
}

func QuickSort(nums []int, left int, right int) {
	// 递归终止条件
	if left >= right {
		return
	}

	// 基准值
	pivot := nums[right]
	// 挡板, 挡板用于保证挡板左边的元素都是比 基准值（pivot）小的
	index := left

	// 这里 i 小于right也可以，right已经被选为基准值了，不需要再取到 right
	for i := left; i <= right-1; i++ {
		if nums[i] < pivot {
			// 小于基准值
			// 移动到左边
			nums[index], nums[i] = nums[i], nums[index]

			// 左挡板加一
			index++
		}
	}
	// 排完一次后，现在挡板的左边 都是比 pivot 小的，右侧元素都是比 pivot 大
	// 所以现在基准值位置 就是 挡板的位置
	nums[index], nums[right] = nums[right], nums[index]

	// 左边部分
	QuickSort(nums, left, index-1)

	//右边部分
	QuickSort(nums, index+1, right)
}
