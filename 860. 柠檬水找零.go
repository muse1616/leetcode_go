package leetcode

func lemonadeChange(bills []int) bool {
	count5 := 0
	count10 := 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			count5++
		} else if bills[i] == 10 {
			if count5 > 0 {
				count5--
				count10++
			} else {
				return false
			}
		} else if bills[i] == 20 {
			if count10 > 0 && count5 > 0 {
				count10--
				count5--
			} else if count10 == 0 && count5 >= 3 {
				count5 -= 3
			}else {
				return false
			}
		}
	}
	return true
}
