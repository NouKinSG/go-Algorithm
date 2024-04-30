package main

// func iceBreaki1ngGame(num int,target int) int{
// 	pre := make([]int,num)
// 	for i := range pre{
// 		pre[i] = i
// 	}

// 	cur := 0
// 	for len(pre) > 1{

// 	}

// }

// LCR 187. 破冰游戏
func iceBreakingGame(num int, target int) int {
	pre := make([]int, num)
	for i := range pre {
		pre[i] = i
	}
	cur := 0
	for len(pre) > 1 {
		cur = (cur + target - 1) % len(pre)
		pre = append(pre[:cur], pre[cur+1:]...)
	}
	return pre[0]
}
