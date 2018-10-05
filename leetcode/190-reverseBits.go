package leetcode

// the process is as follow:
// abcdefgh -> efghabcd -> ghefcdab -> hgfedcba
func reverseBits(n uint32) uint32 {
	if n == 0 {
		return 0
	}
	n = n>>16 | n<<16
	n = (n&0xff00ff00)>>8 | (n&0x00ff00ff)<<8
	n = (n&0xf0f0f0f0)>>4 | (n&0x0f0f0f0f)<<4
	n = (n&0xcccccccc)>>2 | (n&0x33333333)<<2
	n = (n&0xaaaaaaaa)>>1 | (n&0x55555555)<<1
	return n
}

func reverseBitsII(n uint32) uint32 {
	if n == 0 {
		return 0
	}
	var res uint32
	for i := uint32(0); i < 32; i++ {
		res <<= 1
		// res |= ((n >> i) & 1) // the same function as follow
		if n&1 == 1 {
			res++ // set bit 1
		}
		n >>= 1
	}
	return res
}
