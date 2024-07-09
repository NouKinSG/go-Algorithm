package two

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
	var Ndigit []int
	for n > 0 {
		Ndigit = append(Ndigit, n%10)
		n /= 10
	}
	sort.Ints(digits)
	//放入map
	m := make(map[int]struct{})
	for i := range digits {
		m[digits[i]] = struct{}{}
	}

	ansDigits := make([]int, len(Ndigit))
	for i := len(Ndigit) - 1; i >= 0; i-- {
		if i > 0 {
			// 找map，看看也没有
			if _, ok := m[Ndigit[i]]; ok {
				ansDigits[i] = Ndigit[i]
				continue
			}

		}

		// 到这就是 map 中没有找到
		// 那就找 小于这个数最大的
		if d := getMaxDigitLtD(digits, Ndigit[i]); d > 0 {
			ansDigits[i] = d
			break
		}

		// 到这里说明，digits中的数，都无法用,需要回溯
		for ; i < len(Ndigit); i++ {
			ansDigits[i] = 0
			// 找小于中最大的
			if d := getMaxDigitLtD(digits, Ndigit[i]); d > 0 {
				ansDigits[i] = d
				break
			}
			// 如果能来到这，说明最高位都没有能用的数
			// 需要降位
			if i == len(Ndigit)-1 {
				ansDigits = ansDigits[:len(ansDigits)-1]
			}
		}
		break
	}

	// 拼接答案
	ans := 0
	for i := len(ansDigits) - 1; i >= 0; i-- {
		ans = ans * 10
		if ansDigits[i] > 0 {
			ans += ansDigits[i]
			continue
		}
		// 如果自己没有数的，用最高位的
		ans += digits[len(digits)-1]
	}
	return ans
}
