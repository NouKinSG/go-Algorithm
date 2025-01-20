package two

func hammingWeight(n int) int {
	ans := 0
	for n != 0 {
		lowDigits := n & 1
		if lowDigits == 1 {
			ans++
		}
		n = n >> 1
	}
	return ans
}
