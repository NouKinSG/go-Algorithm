package five

import "sort"

// 小于 n 的最大数
func getMaxNumLtN(digits []int, n int) int {
	var Ndigits []int
	for n > 0 {
		Ndigits = append(Ndigits, n%10)
		n /= 10
	}

	// 排序和map
	sort.Ints(digits)
	m := make(map[int]struct{})
	for i := range digits {
		m[digits[i]] = struct{}{}
	}

	ansDigits := make([]int, len(Ndigits)) // TODO  这里的 ansDigits 长度应该和 Ndigits 对齐，因为是小于等于n的最大数
	for i := len(Ndigits) - 1; i >= 0; i-- {
		if i > 0 {
			if _, ok := m[Ndigits[i]]; ok {
				ansDigits[i] = Ndigits[i]
				continue
			}
		}

		if d := getMaxDigitLtD(digits, Ndigits[i]); d > 0 {
			ansDigits[i] = d
			break
		}

		// 回溯
		for ; i < len(Ndigits); i++ {
			ansDigits[i] = 0
			if d := getMaxDigitLtD(digits, Ndigits[i]); d > 0 {
				ansDigits[i] = d
				break
			}
			// 需要降位
			if i == len(Ndigits)-1 {
				ansDigits = ansDigits[:len(ansDigits)-1]
			}
		}
		break
	}

	// 组装答案
	ans := 0
	for i := len(ansDigits) - 1; i >= 0; i-- {
		ans *= 10
		if ansDigits[i] > 0 {
			ans += ansDigits[i]
			continue
		} else {
			ans += digits[len(digits)-1]
		}
	}
	return ans
}

func getMaxDigitLtD(digits []int, digit int) int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < digit {
			return digits[i]
		}
	}
	return 0
}
