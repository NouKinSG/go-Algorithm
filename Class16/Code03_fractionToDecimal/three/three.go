package three

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	/*
		思路：
		1、边界处理
		2、定义答案收集起来
		3、整数部分
		4、小数部分
	*/

	// 边界处理
	if numerator == 0 {
		return ""
	}

	// 定义答案
	var ans []byte
	if (numerator < 0) != (denominator < 0) {
		ans = append(ans, '-')
	}

	// 开始计算结果
	absNum := abs(numerator)
	absDen := abs(denominator)

	// 整数部分
	integrate := absNum / absDen
	ans = append(ans, strconv.Itoa(integrate)...)

	// 小数部分
	// 先找余数
	remainder := absNum % absDen

	if remainder == 0 {
		return string(ans)
	}

	ans = append(ans, '.')
	remainderMap := make(map[int]int)
	for remainder != 0 {
		// 看看有没有余数出现过
		if pos, ok := remainderMap[remainder]; ok {
			return string(ans[:pos]) + "(" + string(ans[pos:]) + ")"
		}

		remainder *= 10
		ans = append(ans, strconv.Itoa(remainder/absDen)...)
		remainder %= absDen
	}
	return string(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
