package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, x int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x)
		a[i] = x
	}

	count := [2]int{1, 0}
	sum, ans := 0, 0
	for i := 0; i < n; i++ {
		sum = (sum + a[i]) % 2
		ans += count[sum]
		count[sum]++
	}
	fmt.Println(ans + 4)
}

//// 第二题
//
//func absInt64(a int64) int64 {
//	if a < 0 {
//		return -a
//	}
//	return a
//}
//
//func main() {
//	var N int
//	fmt.Scan(&N)
//
//	w := make([]int64, N)
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Split(bufio.ScanWords)
//	for i := 0; i < N; i++ {
//		scanner.Scan()
//		fmt.Sscan(scanner.Text(), &w[i])
//	}
//
//	dp := make([][]int64, N)
//	for i := range dp {
//		dp[i] = make([]int64, N)
//	}
//
//	for i := 0; i < N; i++ {
//		dp[i][i] = absInt64(w[i])
//	}
//
//	for length := 2; length <= N; length++ {
//		for i := 0; i <= N-length; i++ {
//			j := i + length - 1
//			for k := i; k < j; k++ {
//				total := absInt64(dp[i][k] + dp[k+1][j])
//
//				if total > dp[i][j] {
//					dp[i][j] = total
//				}
//			}
//		}
//	}
//	fmt.Println(dp[0][N-1])
//}

// 第一题
//func main() {
//	var x, y, t int
//	fmt.Scan(&x, &y, &t)
//
//	for i := 0; i < t; i++ {
//		var r int
//		fmt.Scan(&r)
//		if r == 1 {
//			x, y = y, -x
//		} else if r == 2 {
//			x, y = -y, x
//		}
//		fmt.Println(x, y)
//	}
//}
