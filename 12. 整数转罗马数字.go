package leetcode


func intToRoman(num int) string {
	strs := [][]string{
		[]string{"","I","II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		[]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		[]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		[]string{"", "M", "MM", "MMM"},
	}
	ret := ""
	i := 0
	for num > 0 {
		ret = strs[i][num%10] + ret
		i++
		num /= 10
	}
	return ret
}
