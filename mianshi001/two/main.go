package main

import (
	"fmt"
	"sort"
)

// 定义区间结构体
type Range struct {
	Start, End int32
}

// 定义区间列表结构体
type RangeList struct {
	ranges []Range // 存储区间的切片
}

// Add 方法：向区间列表中添加一个新的区间
// 如果新区间与已有区间重叠或相邻，则合并
func (r *RangeList) Add(rb, re int32) {
	if rb >= re {
		return // 无效区间，结束操作
	}

	newRanges := []Range{} // 存储合并后的区间
	inserted := false      // 标记是否已经插入新区间

	// 遍历现有的区间列表
	for _, curr := range r.ranges {
		if re < curr.Start { // 新区间在当前区间左侧
			if !inserted {
				// 插入新区间
				newRanges = append(newRanges, Range{rb, re})
				inserted = true
			}
			newRanges = append(newRanges, curr) // 当前区间直接加入
		} else if rb > curr.End { // 新区间在当前区间右侧
			newRanges = append(newRanges, curr) // 当前区间直接加入
		} else { // 新区间与当前区间重叠或相邻，进行合并
			rb = min(rb, curr.Start) // 更新新区间的起点
			re = max(re, curr.End)   // 更新新区间的终点
		}
	}

	if !inserted {
		// 如果新区间没有插入，说明新区间是最后一个，直接添加
		newRanges = append(newRanges, Range{rb, re})
	}

	// 更新区间列表
	r.ranges = newRanges
	r.mergeRanges() // 确保所有区间排序并去重叠
}

// Remove 方法：从区间列表中移除一个区间
// 如果移除区间与现有区间有交集，进行分割
func (r *RangeList) Remove(rb, re int32) {
	if rb >= re {
		return // 无效区间，结束操作
	}

	newRanges := []Range{} // 存储修改后的区间列表

	// 遍历现有区间列表
	for _, curr := range r.ranges {
		if re <= curr.Start || rb >= curr.End { // 没有交集，直接保留当前区间
			newRanges = append(newRanges, curr)
		} else { // 有交集，需要分割当前区间
			if curr.Start < rb {
				// 当前区间左侧部分保留
				newRanges = append(newRanges, Range{curr.Start, rb})
			}
			if curr.End > re {
				// 当前区间右侧部分保留
				newRanges = append(newRanges, Range{re, curr.End})
			}
		}
	}

	// 更新区间列表
	r.ranges = newRanges
}

// ToString 方法：将区间列表转换为字符串表示
func (r *RangeList) ToString() string {
	result := "" // 存储最终结果的字符串
	// 遍历所有区间
	for _, rg := range r.ranges {
		result += fmt.Sprintf("[%d, %d) ", rg.Start, rg.End) // 拼接每个区间的字符串表示
	}
	return result[:len(result)-1] // 去掉最后的空格
}

// mergeRanges 方法：确保所有区间按起始位置排序，并且无重叠
func (r *RangeList) mergeRanges() {
	// 按照区间的起始位置进行排序
	sort.Slice(r.ranges, func(i, j int) bool {
		return r.ranges[i].Start < r.ranges[j].Start
	})

	merged := []Range{} // 存储合并后的区间
	// 遍历区间列表，进行合并
	for _, curr := range r.ranges {
		// 如果merged为空或当前区间与merged最后一个区间不重叠，直接添加
		if len(merged) == 0 || merged[len(merged)-1].End < curr.Start {
			merged = append(merged, curr)
		} else {
			// 如果重叠，合并区间
			merged[len(merged)-1].End = max(merged[len(merged)-1].End, curr.End)
		}
	}

	// 更新区间列表
	r.ranges = merged
}

// 辅助函数：返回两个数中的较小值
func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// 辅助函数：返回两个数中的较大值
func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// 主函数：测试 RangeList 的功能
func main() {
	r := &RangeList{}

	// 插入区间 [5, 10)
	r.Add(5, 10)
	fmt.Println(r.ToString()) // [5, 10)

	// 插入区间 [1, 3)，插入后区间顺序不正确
	r.Add(1, 3)
	fmt.Println(r.ToString()) // [5, 10) [1, 3)

	// 插入区间 [2, 6)，插入后存在重叠，但没有进行合并
	r.Add(2, 6)
	fmt.Println(r.ToString()) // [5, 10) [1, 3) [2, 6)

	// 通过 mergeRanges() 来合并并确保按顺序排列
	r.mergeRanges()
	fmt.Println(r.ToString()) // [1, 10)
}

//func main() {
//	r := &RangeList{}
//
//	r.Add(1, 5)
//	fmt.Println(r.ToString()) // 应输出: [1, 5)
//
//	r.Add(10, 20)
//	fmt.Println(r.ToString()) // 应输出: [1, 5) [10, 20)
//
//	r.Add(20, 20)
//	fmt.Println(r.ToString()) // 应输出: [1, 5) [10, 20)
//
//	r.Add(20, 21)
//	fmt.Println(r.ToString()) // 应输出: [1, 5) [10, 21)
//
//	r.Add(2, 4)
//	fmt.Println(r.ToString()) // 应输出: [1, 5) [10, 21)
//
//	r.Add(3, 8)
//	fmt.Println(r.ToString()) // 应输出: [1, 8) [10, 21)
//
//	r.Remove(10, 10)
//	fmt.Println(r.ToString()) // 应输出: [1, 8) [10, 21)
//
//	r.Remove(10, 11)
//	fmt.Println(r.ToString()) // 应输出: [1, 8) [11, 21)
//
//	r.Remove(15, 17)
//	fmt.Println(r.ToString()) // 应输出: [1, 8) [11, 15) [17, 21)
//
//	r.Remove(3, 19)
//	fmt.Println(r.ToString()) // 应输出: [1, 3) [19, 21)
//}
