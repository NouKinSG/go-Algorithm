package test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestOS(t *testing.T) {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	} // 读取前两个整数
	fmt.Println(a + b) // 输出他们的和
}

func TestBUFIO(t *testing.T) {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {

		}
	}(writer)

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		nums := strings.Split(line, " ")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		fmt.Fprintln(writer, a+b)
	}
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		if j, ok := hashMap[target-num]; ok {
			return []int{i, j}
		}
		hashMap[num] = i
	}
	return nil
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//func longestPalindrome(s string) string {
//	if len(s) < 2 {
//		return s
//	}
//	var ans string
//	// 头尾指针
//	start,end := 0,0
//
//	for i := 0; i < len(s); i++{
//		len1 :=
//	}
//
//	return ans
//}
//
//// 中心扩散法
//func expandAroundCenter(s string,left,right int) int{
//	for left >= 0 && right < len(s) && s[left] == s[right]{
//		left--
//		right++
//	}
//	return right - left - 1
//}
