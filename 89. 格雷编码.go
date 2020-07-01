package leetcode

import "math"

func grayCode(n int) []int {

	//如 n = 3:
	//G(i) = i ^ (i/2)
	//G(0) = 000,
	//G(1) = 1 ^ 0 = 001 ^ 000 = 001
	//G(2) = 2 ^ 1 = 010 ^ 001 = 011
	//G(3) = 3 ^ 1 = 011 ^ 000 = 010
	//G(4) = 4 ^ 2 = 100 ^ 010 = 110
	//G(5) = 5 ^ 2 = 101 ^ 010 = 111
	//G(6) = 6 ^ 3 = 110 ^ 011 = 101
	//G(7) = 7 ^ 3 = 111 ^ 011 = 100

	ans := make([]int, 0)

	for i := 0; i < int(math.Pow(float64(2), float64(n))); i++ {
		if i == 0 {
			ans = append(ans, 0)
		} else {
			ans = append(ans, i^(i/2))
		}
	}

	return ans

	//res := make([]int, 1<<n)
	//index := 1
	//for i := 0; i < n; i += 1 {
	//	r := 1<<i - 1
	//	for r >= 0 {
	//		v := res[r]
	//		v = v | (1 << i)
	//		res[index] = v
	//		r -= 1
	//		index += 1
	//	}
	//}
	//return res
}
