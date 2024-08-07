package one

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	// 1、边界处理：0、负数
	// 2、定义答案、收集结果
	// 3、收集整数部分结果
	// 4、收集小数部分结果  eg.小数部分需要计算余数

	if numerator == 0 {
		return "0"
	}

	var result []byte
	// 判断符号
	if (numerator < 0) != (denominator < 0) {
		result = append(result, '-')
	}

	//使用绝对值来处理计算
	absNum := abs(numerator)
	absDen := abs(denominator)

	//整数部分
	integerPart := absNum / absDen
	result = append(result, strconv.Itoa(integerPart)...)

	// 开始寻找小数部分，需要余数

	// 寻找余数
	remainder := absNum % absDen
	if remainder == 0 {
		// 余数为0，说明没有小数部分
		return string(result)
	}

	result = append(result, '.')      // 先加小数点
	remainderMap := make(map[int]int) // 用一个余数map，收集每一位余数，如果再次出现说明有循环

	for remainder != 0 {
		if pos, ok := remainderMap[remainder]; ok {
			// 发现余数再次出现，说明是循环
			return string(result[:pos]) + "(" + string(result[pos:]) + ")"
		}

		// 记录当前余数的位置
		remainderMap[remainder] = len(result)

		// 计算下一位小数
		remainder *= 10
		result = append(result, strconv.Itoa(remainder/absDen)...)
		remainder %= absDen
	}
	return string(result)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
