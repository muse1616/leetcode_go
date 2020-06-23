package leetcode


func reverseBits(num uint32) uint32 {
	result := uint32(0)
	pow := uint32(31)
	for num != 0 {
		result += (num & 1) << pow
		num = num >> 1
		pow --
	}
	return result
}
