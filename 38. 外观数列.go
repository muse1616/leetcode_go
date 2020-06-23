package leetcode

import "fmt"

func countAndSay(n int) string {
	str := "1"
	if n == 1 {
		return str
	}
	for j := 2; j <= n; j++ {
		l, tmp := len(str), str
		t, c := tmp[0], 1
		str = ""

		for i := 1; i < l; i++ {
			if tmp[i] == t {
				c++
			} else {
				str += fmt.Sprintf("%d%d", c, t-'0')
				c = 1
				t = tmp[i]
			}
		}
		str += fmt.Sprintf("%d%d", c, t-'0')
	}
	return str
}
