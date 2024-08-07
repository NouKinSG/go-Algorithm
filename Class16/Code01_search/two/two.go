package two

/*
理解问题：
1、你有一个升序排列但被旋转过的数组。
2、需要在这个数组中找到目标值的下标，如果不存在，返回 -1。
3、旋转数组中没有重复元素。
*/

/*
分析特性：
	1、旋转后的数组可以被分为两部分：一个是有序的，一个是旋转过的无序部分。
	2、例如：在 [4,5,6,7,0,1,2] 中，[4,5,6,7] 是有序的，[0,1,2] 也是有序的，但因为旋转，使整体无序。
*/

/*
步骤详解：

初始化：

	1、设置两个指针 left 和 right，分别指向数组的起始和终止位置。

循环查找：
	1、使用 while left <= right 作为循环条件。
	2、计算中间位置：mid = (left + right) // 2。
	3、检查中间元素是否是目标值：if nums[mid] == target: return mid。

判断有序部分：

	左半部分有序：
		1、如果 nums[left] <= nums[mid]，说明 left 到 mid 是有序的。
		2、检查目标值是否在 left 和 mid 之间：
		3、如果 nums[left] <= target < nums[mid]，目标在左半部分，更新 right = mid - 1。
		4、否则，目标在右半部分，更新 left = mid + 1。

	右半部分有序：
		1、如果 nums[mid] < nums[right]，说明 mid 到 right 是有序的。
		2、检查目标值是否在 mid 和 right 之间：
		3、如果 nums[mid] < target <= nums[right]，目标在右半部分，更新 left = mid + 1。
		4、否则，目标在左半部分，更新 right = mid - 1。

未找到目标值：
	如果退出循环而没有找到目标值，则返回 -1。

*/

func search(nums []int, target int) int {
	// 初始化
	left, right := 0, len(nums)-1

	// 开找
	for left <= right {
		// 中间位置
		mid := left + (right-left)>>1

		// 想用二分查找
		if nums[mid] == target {
			return mid
		}

		// 找有序部分
		// 因为数组被旋转了，但还是局部有序，整体无序
		// 也就是说，如果 nums[left] 小于或者等于 nums[mid] 那就表明了，left 到 mid 这局部是有序的
		if nums[left] <= nums[mid] {
			// 看看目标值target 在不在这个有序区间
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 对于 else 的 右半部分有序 解释
			// 这里可能会有点疑惑。 else 不是等于 nums[left] > nums[mid] 吗，为什么这个相当于右半部分有序？
			// 首先，原数组是升序的，旋转了，也是局部升序。 5 6 7 0 1 2 3 这样子。
			// 如果我们的  left > mid 说明右侧一定是升序的。 比如 这里  5 > 0 说明 0(mid) 到 3(right) 有序

			// 看看目标值target 在不在这个有序区间
			if nums[mid] <= target && target < nums[right] {
				// 在这个区间内
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	// 没找到
	return -1
}
