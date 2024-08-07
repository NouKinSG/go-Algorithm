package six

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

	ansDigits := make([]int, len(Ndigits))
	// 高到低位开始遍历
	for i := len(Ndigits) - 1; i >= 0; i-- {
		if i > 0 {
			if _, ok := m[Ndigits[i]]; ok {
				ansDigits[i] = Ndigits[i]
				continue
			}
		}

		// map中没有，用小于的最大
		if d := getMaxNumLtDigit(digits, Ndigits[i]); d > 0 {
			ansDigits[i] = d
			break
		}

		// 如果还是没有，回溯
		for ; i < len(Ndigits); i++ {
			ansDigits[i] = 0
			// 再找一次
			if d := getMaxNumLtDigit(digits, Ndigits[i]); d > 0 {
				ansDigits[i] = d
				break
			}
			// 如果还是没有
			// 降位
			if i == len(Ndigits)-1 {
				ansDigits = ansDigits[:len(ansDigits)-1]
			}
		}
		break
	}

	ans := 0
	// 拼接答案
	for i := len(ansDigits) - 1; i >= 0; i-- {
		ans *= 10
		if ansDigits[i] > 0 {
			ans += ansDigits[i]
		} else {
			// 用digits的最高位
			ans += digits[len(digits)-1]
		}
	}
	return ans
}

// 有序数组找小于digit的最大数
func getMaxNumLtDigit(digits []int, digit int) int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < digit {
			return digits[i]
		}
	}
	return 0
}
