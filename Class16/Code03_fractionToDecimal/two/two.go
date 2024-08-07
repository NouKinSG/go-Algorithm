package two

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {

	// 1、边界处理: 负数、0
	if numerator == 0 {
		return "0"
	}

	// 2、定义答案，用于收集结果
	var ans []byte
	// 负数情况
	if (numerator < 0) != (denominator < 0) {
		ans = append(ans, '-')
	}

	// 计算的结果用 绝对值去计算
	absNum := abs(numerator)
	absDen := abs(denominator)

	// 3、整数部分
	integerPart := absNum / absDen
	ans = append(ans, strconv.Itoa(integerPart)...)

	//  先找余数
	remainder := absNum % absDen
	// 如果余数是 0，就没有小数部分了
	if remainder == 0 {
		return string(ans)
	}

	// 4、小数部分

	// 4.1 加小数点
	ans = append(ans, '.')
	// 4.2 用 map 收集余数，如果再次出现就是循环
	remainderMap := make(map[int]int)

	for remainder != 0 {
		// 看看这个余数有没有出现过，如果出现过说明是循环小数
		if pos, ok := remainderMap[remainder]; ok {
			return string(ans[:pos]) + "(" + string(ans[pos:]) + ")"
		}

		// 记录这个余数的位置
		remainderMap[remainder] = len(ans)

		// 计算下一位
		remainder *= 10
		// 用余数来除，才能获得小数
		ans = append(ans, strconv.Itoa(remainder/absDen)...)
		remainder %= absDen
	}
	return string(ans)
}

func abs(a int) int {
	if a < 0 {
		return a
	}
	return a
}
