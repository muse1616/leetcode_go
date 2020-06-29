package leetcode

func hasGroupsSizeX(deck []int) bool {
	//长度为 0
	if len(deck) == 0 {
		return false
	}

	hash := make(map[int]int)

	for _, v := range deck {
		hash[v]++
	}
	maxVal := hash[deck[0]]
	for _, v := range hash {
		//求最大公约数
		maxVal = gcd(v, maxVal)
		if maxVal < 2 {
			return false
		}
	}
	return true
}
//辗转相除法
func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}
