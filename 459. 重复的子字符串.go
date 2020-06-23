package leetcode

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}
	b := []byte(s)
	//依次从开头比较长度为1/2/3/4---len/2-1的子串能否构成字符串
	for i := 0; i <= len(b)/2-1; i++ {
		//当前字符
		tmp := b[:i+1]
		//跳出二次循环标志
		f := false
		for j := i + 1; j+len(tmp) <= len(b); j += len(tmp) {
			m:=j
			for k := 0; k < len(tmp); k++ {
				if tmp[k] != b[m] {
					f = true
					break
				}
				m++
			}
			if f == true {
				break
			}
			if m-1==len(b)-1{
				return true
			}
		}
	}
	return false
}
