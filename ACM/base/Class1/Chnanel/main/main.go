package main

import (
	"fmt"
	"sort"
	"strconv"
)

// getMaxDigitLtD 获取小于指定数字的最大数字。
func getMaxDigitLtD(digits []int, digit int) int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < digit {
			return digits[i]
		}
	}
	return -1 // 如果没有小于指定数字的数字，返回 -1
}

// findMaxNumLtN 获取小于 n 的最大数。
func findMaxNumLtN(digits []int, n int) int {
	nStr := strconv.Itoa(n) // 将 n 转换为字符串形式，方便逐位处理
	sort.Ints(digits)       // 对 digits 进行排序，方便查找和比较
	m := make(map[int]struct{})
	for _, v := range digits {
		m[v] = struct{}{}
	}

	// 用于构造结果的切片
	result := make([]byte, len(nStr))
	// 标记是否已找到小于对应位的数字
	foundSmaller := false

	for i := 0; i < len(nStr); i++ {
		curDigit := int(nStr[i] - '0')

		// 如果已找到小于当前位的数字，直接使用 digits 中的最大值
		if foundSmaller {
			result[i] = byte(digits[len(digits)-1] + '0')
			continue
		}

		// 查找当前位小于 curDigit 的最大数字
		maxDigit := getMaxDigitLtD(digits, curDigit)
		if maxDigit != -1 {
			result[i] = byte(maxDigit + '0')
			foundSmaller = true
		} else if _, exists := m[curDigit]; exists {
			// 如果 digits 中存在与当前位相同的数字，使用该数字
			result[i] = byte(curDigit + '0')
		} else {
			// 回溯处理，找到可以替换的位
			for j := i - 1; j >= 0; j-- {
				origDigit := int(result[j] - '0')
				newDigit := getMaxDigitLtD(digits, origDigit)
				if newDigit != -1 {
					result[j] = byte(newDigit + '0')
					for k := j + 1; k < len(nStr); k++ {
						result[k] = byte(digits[len(digits)-1] + '0')
					}
					foundSmaller = true
					break
				}
			}
			if !foundSmaller {
				return -1 // 如果最高位也没有找到合适的替换，返回 -1
			}
		}
	}

	// 将构造的结果转换为整数
	finalResult, _ := strconv.Atoi(string(result))
	return finalResult
}

func main() {
	n := 23121
	A := []int{2, 4, 9}
	result := findMaxNumLtN(A, n)
	fmt.Printf("由 %v 组成的小于 %d 的最大数是: %d\n", A, n, result)
}
