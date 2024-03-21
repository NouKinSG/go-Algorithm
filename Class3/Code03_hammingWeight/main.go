package main

func hammingWeight(num uint32) int {
	ans := 0
	for i := 0; i < 32; i++ {
		if num&(1<<i) > 0 {
			ans++
		}
	}
	return ans
}
