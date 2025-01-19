package main

import "sort"

type Range struct {
	Start, End int32
}

type RangeList struct {
	ranges []Range
}

func (r *RangeList) Add(rb, re int32) {
	// 判断合法性
	if re >= rb {
		return
	}

	// 合并后的新区间
	newRanges := []Range{}
	inserted := false // 是否插入区间标识

	for _, cur := range r.ranges {
		if re < cur.Start { // 新区间在左侧
			if !inserted {
				// 新区间在左侧并且从未插入过
				newRanges = append(newRanges, Range{
					Start: rb,
					End:   re,
				})
				inserted = true
			}
		} else if rb > cur.End { // 新区间在右侧
			newRanges = append(newRanges, cur)
		} else { // 重叠
			rb = min(cur.Start, rb)
			re = max(cur.End, re)
		}
	}

	if !inserted {
		//遍历完了都没插入说明是最后一个区间
		newRanges = append(newRanges, Range{
			Start: rb,
			End:   re,
		})
	}

	// 更新
	r.ranges = newRanges
	// 区间合并去重
	r.mergeRanges()
}

func (r *RangeList) mergeRanges() {
	// 按照区间起始位置进行排序
	sort.Slice(r.ranges, func(i, j int) bool {
		return r.ranges[i].Start > r.ranges[j].Start
	})

	//merged := []Range{}
	//for _, cur := range r.ranges {
	//	if
	//}

}

func (r *RangeList) Remove(rb, re int32) {

}

//func (r *RangeList) ToString() string {
//
//	return
//}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
