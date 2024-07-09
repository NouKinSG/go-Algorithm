package one

import "sort"

// getMaxNumLtN 获取小于 n 的最大数。
func getMaxNumLTN(digits []int, n int) int {
	// 1. 将 n 转换为数组，方便高位到低位操作
	// 2. 排序，放入 map 方便取
	// 3. 从高位到低位遍历，尽量取相同值
	// 4. 回溯
	// 5. 拼接成目标数

	// 1. 将 n 转换为数组，方便高位到低位操作
	var ndigits []int
	for n > 0 {
		ndigits = append(ndigits, n%10)
		n /= 10
	}

	// 2.1 排序给定数组
	sort.Ints(digits)
	// 2.2 存入map
	m := make(map[int]struct{})
	for i := range digits {
		m[digits[i]] = struct{}{}
	}

	// 3. 从高位开始遍历了
	ansArr := make([]int, len(ndigits))
	for i := len(ndigits) - 1; i >= 0; i-- {
		if i > 0 {
			// 没到最后一位的时候
			//如果存在相同数字
			if _, ok := m[ndigits[i]]; ok {
				ansArr[i] = ndigits[i]
				continue
			}

			// 能到这里说明不存在相同数字
			// 如果不存在相同数字，就找小的里面最大的
			if d := getMaxDigitLtD(digits, ndigits[i]); d > 0 {
				ansArr[i] = d
				break
			}
		}
		if i == 0 {
			// 最后一位
			if d := getMaxDigitLtD(digits, ndigits[i]); d > 0 {
				ansArr[i] = d
				break
			}
		}

		// 4. 回溯
		for ; i < len(ndigits); i++ {
			ansArr[i] = 0
			if d := getMaxDigitLtD(digits, ndigits[i]); d > 0 {
				ansArr[i] = d
				break
			}
			// 最高位也没有小于其的最大数字
			if i == len(ndigits)-1 {
				// 需要 降位了
				ansArr = ansArr[:len(ansArr)-1]
			}
		}
		break
	}

	// 5. 拼接目标数
	var ans int
	for i := len(ansArr) - 1; i >= 0; i-- {
		ans = ans * 10
		if ansArr[i] > 0 {
			// 如果 某一位不是 0 ，就直接有自己的数
			ans += ansArr[i]
			continue
		}
		// 如果没有自己的数，就直接用最大的一位
		ans += digits[len(digits)-1]
	}

	return ans
}

// 从排序好的数组中找小于 digit 最大的数
func getMaxDigitLtD(digits []int, digit int) int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < digit {
			return digits[i]
		}
	}
	return 0
}
