package leetcode

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A==B{
		hash:=make(map[byte]int)
		for i:=0;i<len(A);i++{
			hash[A[i]]++
			if hash[A[i]] == 2{
				return true
			}
		}
	}
	var a []byte
	var b []byte
	for i:=0;i<len(A);i++{
		if A[i]!=B[i]{
			a = append(a,A[i])
			b = append(b,B[i])
		}
	}
	return len(a)==2&&len(b)==2&&a[0]==b[1]&&a[1]==b[0]

}
