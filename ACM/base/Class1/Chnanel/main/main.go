package main

import (
	"fmt"
	"sort"
)

// getMaxDigitLtD 获取小于指定数字的数字。
func getMaxDigitLtD(digits []int, digit int) int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < digit {
			return digits[i]
		}
	}
	return 0
}

// getMaxNumLtN 获取小于 n 的最大数。
func getMaxNumLTN(digits []int, n int) int {
	var ndigits []int
	// 获取 n 的每一位数字。
	for n > 0 {
		ndigits = append(ndigits, n%10)
		n /= 10
	}
	// 排序给定的数字数组。
	sort.Ints(digits)

	// 数字写入 map 用于查看是否存在。
	m := make(map[int]struct{})
	for _, v := range digits {
		m[v] = struct{}{}
	}
	// 目标数的每一位数字。
	tdigits := make([]int, len(ndigits))

	// 从高位遍历，尽可能地取相同值，除了最后一位。
	for i := len(ndigits) - 1; i >= 0; i-- {
		if i > 0 {
			// 存在相同数字。
			if _, ok := m[ndigits[i]]; ok {
				tdigits[i] = ndigits[i]
				continue
			}
			// 存在小于当前位的最大数字。
			if d := getMaxDigitLtD(digits, ndigits[i]); d > 0 {
				tdigits[i] = d
				break
			}
		}
		if i == 0 {
			// 存在小于当前数字的最大数字。
			if d := getMaxDigitLtD(digits, ndigits[i]); d > 0 {
				tdigits[i] = d
				break
			}
		}
		// 回溯
		for ; i < len(ndigits); i++ {
			tdigits[i] = 0
			if d := getMaxDigitLtD(digits, ndigits[i]); d > 0 {
				tdigits[i] = d
				break
			}
			// 最高位也没有小于其的最大数字。
			if i == len(ndigits)-1 {
				tdigits = tdigits[:len(tdigits)-1]
			}
		}
		break
	}
	// 拼接目标数。
	var target int
	for i := len(tdigits) - 1; i >= 0; i-- {
		target *= 10
		if tdigits[i] > 0 {
			target += tdigits[i]
			continue
		}
		target += digits[len(digits)-1]
	}
	return target
}

func main() {
	fmt.Println(getMaxNumLTN([]int{1, 2, 9, 4}, 2533))
	fmt.Println(getMaxNumLTN([]int{1, 2, 5, 4}, 2543))
	fmt.Println(getMaxNumLTN([]int{1, 2, 5, 4}, 2541))
	fmt.Println(getMaxNumLTN([]int{1, 2, 9, 4}, 2111))
	fmt.Println(getMaxNumLTN([]int{5, 9}, 5555))
	fmt.Println(getMaxNumLTN([]int{3, 5, 4}, 2355))
	fmt.Println(getMaxNumLTN([]int{3, 5, 4}, 23))
	fmt.Println(getMaxNumLTN([]int{3, 5, 4}, 2555))

}
