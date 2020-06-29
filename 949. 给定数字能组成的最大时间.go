package leetcode

import "fmt"

func largestTimeFromDigits(A []int) string {
	ans := -1
	hour, min := 0, 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j != i {
				for k := 0; k < 4; k++ {
					if k != i && k != j {
						v := 6 - i - j - k
						hour = A[i]*10 + A[j]
						min = A[k]*10 + A[v]
						if hour < 24 && min < 60 {
							if hour*60+min > ans {
								ans = hour*60 + min
							}
						}
					}
				}
			}
		}
	}
	if ans >= 0 {
		return fmt.Sprintf("%02d:%02d", ans/60, ans%60)
	}
	return ""

}
