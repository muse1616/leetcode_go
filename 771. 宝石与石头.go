package leetcode

func numJewelsInStones(J string, S string) int {
	hash := make(map[byte]bool, 10)
	for _, v := range J {
		hash[byte(v)] = true
	}
	sum:=0
	for _,v:=range S{
		if _,ok:=hash[byte(v)];ok{
			sum++
		}

	}
	return sum

}
