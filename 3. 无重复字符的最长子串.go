package leetcode

func lengthOfLongestSubstring(s string) int {
	//已遍历序号
	i := 0
	//最大长度
	max := 0
	//转为rune切片
	a := []rune(s)
	for m, c := range a {
		for n := i; n < m; n++ {
			//出现重复字母
			if a[n] == c {
				//i直接跳到n的后一个字符 即第一次出现重复字母的后一个位置 此时前方已无需比较
				i = n + 1
				break
			}
		}
		//若此长度大于max重新赋值max
		if m-i+1 > max {
			max = m - i + 1
		}
	}

	return max
}
