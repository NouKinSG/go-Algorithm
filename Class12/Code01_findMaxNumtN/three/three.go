package three

import "sort"

func getMaxDigitLtD(digits []int, digit int) int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < digit {
			return digits[i]
		}
	}
	return 0
}

func getMaxNumLtN(digits []int, n int) int {
	var Ndigits []int
	for n > 0 {
		Ndigits = append(Ndigits, n%10)
		n /= 10
	}

	sort.Ints(digits)
	// 加入map
	m := make(map[int]struct{})
	for i := range digits {
		m[digits[i]] = struct{}{}
	}

	ansDigtis := make([]int, len(Ndigits))
	for i := len(Ndigits) - 1; i >= 0; i-- {
		if i > 0 {

			if _, ok := m[Ndigits[i]]; ok {
				ansDigtis[i] = Ndigits[i]
				continue
			}
		}

		if d := getMaxDigitLtD(digits, Ndigits[i]); d > 0 {
			ansDigtis[i] = d
			break
		}

		// 回溯
		for ; i < len(Ndigits); i++ {
			ansDigtis[i] = 0

			if d := getMaxDigitLtD(digits, Ndigits[i]); d > 0 {
				ansDigtis[i] = d
				break
			}

			// 如果还不行，降位处理
			if i == len(Ndigits)-1 {
				ansDigtis = ansDigtis[:len(ansDigtis)-1]
			}
		}
		break
	}

	// 拼接答案
	ans := 0
	for i := len(ansDigtis) - 1; i >= 0; i-- {
		ans *= 10
		if ansDigtis[i] > 0 {
			ans += ansDigtis[i]
			continue
		}
		ans += digits[len(digits)-1]
	}

	return ans
}
