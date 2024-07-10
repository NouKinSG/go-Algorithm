package four

import "sort"

// 小于 n 的最大数
func getMaxNumLtN(digits []int, n int) int {
	//1、 n 变数组方便从高位开始
	var Ndigits []int
	for n > 0 {
		Ndigits = append(Ndigits, n%10)
		n /= 10
	}
	//2、 排序 和 map
	sort.Ints(digits)
	m := make(map[int]struct{})
	for i := range digits {
		m[digits[i]] = struct{}{}
	}

	// 开始从高位遍历
	ansDigits := make([]int, len(Ndigits))
	for i := len(Ndigits) - 1; i >= 0; i-- {
		if i > 0 {
			// 有无相同
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

			//降位
			if i == len(Ndigits)-1 {
				ansDigits = ansDigits[:len(digits)-1]
			}
		}
		break
	}

	// 拼接答案
	ans := 0
	for i := len(ansDigits) - 1; i >= 0; i-- {
		ans *= 10
		if ansDigits[i] > 0 {
			ans += ansDigits[i]
		}
		// 没有
		ans += digits[len(digits)-1]
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
