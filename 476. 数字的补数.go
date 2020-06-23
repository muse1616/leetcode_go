package leetcode
func findComplement(num int) int {
	//x+x补数 == 2^N -1
	i:=1
	for i<=num{
		i <<= 1
	}
	return i-num-1
}