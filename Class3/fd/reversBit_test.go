package fd_test

import "testing"

func reverseBits(num uint32) uint32 {
	var ans uint32 = 0
	for i := 0; i < 32; i++ {
		var a uint32 = num & (1 << i)
		ans = ans + (a << (31 - i))
		// num = num >> 1
	}
	return ans
}

func TestReverBit(t *testing.T) {
	var temp uint32 = uint32(0b11111111111111111111111111111110)
	ans := reverseBits(temp)
	t.Log(ans)
}
