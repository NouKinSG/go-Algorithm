package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(n int, a []int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		cnt := make([]int, n+2)
		vis := make([]bool, n+2)
		mex := 0
		for j := i; j >= 1; j-- {
			cnt[a[j-1]]++
			vis[a[j-1]] = true
			for vis[mex] {
				mex++
			}
			dp[i] = max(dp[i], dp[j-1]+mex)
		}
	}
	return dp[n]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	results := make([]int, 0, T)

	for T > 0 {
		T--

		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		a := make([]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			a[i], _ = strconv.Atoi(scanner.Text())
		}
		results = append(results, solve(n, a))
	}

	for _, results := range results {
		fmt.Println(results)
	}
}

//func solve() {
//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//
//	var T int64
//	fmt.Fscan(in, &T)
//	for ; T > 0; T-- {
//		var n, k int64
//		var s string
//		fmt.Fscan(in, &n, &k, &s)
//		arr := []byte(s)
//		i, j := int64(0), int64(n-1)
//
//		for k > int64(0) && i < j {
//			for i < n && arr[i] == '0' {
//				i++
//			}
//			for j >= 0 && arr[j] == '1' {
//				j--
//			}
//			if i >= j {
//				break
//			}
//			arr[i], arr[j] = arr[j], arr[i]
//			k--
//			i++
//			j--
//		}
//		fmt.Fprintln(out, string(arr))
//	}
//}
//
//func main() {
//	solve()
//}

//var digits = [10][5][4]bool{
//	// 0
//	{
//		{true, true, true, true},
//		{true, false, false, true},
//		{true, false, false, true},
//		{true, false, false, true},
//		{true, true, true, true},
//	},
//	// 1
//	{
//		{false, false, true, true},
//		{false, false, false, true},
//		{false, false, false, true},
//		{false, false, false, true},
//		{false, false, false, true},
//	},
//	// 2
//	{
//		{true, true, true, true},
//		{false, false, false, true},
//		{true, true, true, true},
//		{true, false, false, false},
//		{true, true, true, true},
//	},
//	// 3
//	{
//		{true, true, true, true},
//		{false, false, false, true},
//		{true, true, true, true},
//		{false, false, false, true},
//		{true, true, true, true},
//	},
//	// 4
//	{
//		{true, false, true, false},
//		{true, false, true, false},
//		{true, true, true, true},
//		{false, false, true, false},
//		{false, false, true, false},
//	},
//	// 5
//	{
//		{true, true, true, true},
//		{true, false, false, false},
//		{true, true, true, true},
//		{false, false, false, true},
//		{true, true, true, true},
//	},
//	// 6
//	{
//		{true, true, true, true},
//		{true, false, false, false},
//		{true, true, true, true},
//		{true, false, false, true},
//		{true, true, true, true},
//	},
//	// 7
//	{
//		{true, true, true, true},
//		{false, false, false, true},
//		{false, false, false, true},
//		{false, false, false, true},
//		{false, false, false, true},
//	},
//	// 8
//	{
//		{true, true, true, true},
//		{true, false, false, true},
//		{true, true, true, true},
//		{true, false, false, true},
//		{true, true, true, true},
//	},
//	// 9
//	{
//		{true, true, true, true},
//		{true, false, false, true},
//		{true, true, true, true},
//		{false, false, false, true},
//		{true, true, true, true},
//	},
//}
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	writer := bufio.NewWriter(os.Stdout)
//	defer writer.Flush()
//
//	var T int
//	fmt.Fscanln(reader, &T)
//	for i := 0; i < T; i++ {
//		var n, x, y int
//		fmt.Fscanln(reader, &n, &x, &y)
//		x--
//		y--
//		s, _ := reader.ReadString('\n')
//		s = strings.TrimSpace(s)
//
//		count := 0
//		currentState := digits[s[0]-'0'][x][y]
//		for j := 1; j < n; j++ {
//			nextState := digits[s[j]-'0'][x][y]
//			if currentState != nextState {
//				count++
//			}
//			currentState = nextState
//		}
//		fmt.Fprintln(writer, count)
//	}
//}
