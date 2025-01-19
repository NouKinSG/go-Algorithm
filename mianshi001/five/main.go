package five

type Range struct {
	Start, End int32
}

type RangeList struct {
	ranges []Range
}

func (r *RangeList) Add(rb, re int32) {
	if re >= rb {
		return
	}
	newRanges := []Range{}
	inserted := false

	for _, cur := range r.ranges {
		if re < cur.Start {
			if !inserted {
				newRanges = append(newRanges, Range{rb, re})
				inserted = true
			}
			newRanges = append(newRanges, cur)
		} else if rb > cur.End {
			newRanges = append(newRanges, cur)
		} else {
			rb = min(rb, cur.Start)
			re = max(re, cur.End)
		}
	}

	if !inserted {
		newRanges = append(newRanges, Range{rb, re})
	}

	r.ranges = newRanges
}

func (r *RangeList) remove(rb, re int32) {
	if re >= rb {
		return
	}
	newRange := []Range{}
	for _, cur := range r.ranges {
		if re <= cur.Start || rb >= cur.End {
			newRange = append(newRange, Range{rb, re})
		} else {
			if cur.Start < rb {
				newRange = append(newRange, Range{cur.Start, rb})
			}
			if cur.End > re {
				newRange = append(newRange, Range{
					Start: re,
					End:   cur.End,
				})
			}
		}
	}
	r.ranges = newRange
}
