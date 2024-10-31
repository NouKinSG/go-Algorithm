package one

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	ans := []string{}
	if len(digits) == 0 {
		return ans
	}

	ans = back("", 0, digits)
	return ans
}

func back(cur string, index int, digits string) []string {
	// 如果已经生成的组合长度等于输入数字字符串长度，说明已完成一个完整组合
	if index == len(digits) {
		return []string{cur}
	}

	// 获取字母映射
	letters := phoneMap[string(digits[index])]

	// 开始组合
	var combinations []string
	for i := 0; i < len(letters); i++ {

		// 用变量接收递归的结果，再进行 append 操作
		nextCombinations := back(cur+string(letters[i]), index+1, digits)

		// 递归调用，在当前位置添加当前字母后继续生成下一个字母
		combinations = append(combinations, nextCombinations...)
	}
	return combinations
}
