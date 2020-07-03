package leetcode

// Reverse bits of a given 32 bits unsigned integer.

// Example 1:
// Input: 00000010100101000001111010011100
// Output: 00111001011110000010100101000000
// Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, so return 964176192 which its binary representation is 00111001011110000010100101000000.

func reverseBits(num uint32) uint32 {
	if num == 0 {
		return 0
	}
	var res uint32
	pos := uint32(31)
	for num != 0 {
		// reverse bit by position
		res |= (num & 1) << pos

		num >>= 1

		pos--
	}
	return res
}

// the process is as follow:
// abcdefgh -> efghabcd -> ghefcdab -> hgfedcba
func reverseBitsII(n uint32) uint32 {
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
