package leetcode



func toLowerCase(str string) string {
	b:=[]byte(str)

	for i:=0;i<len(b);i++{
		if b[i]<=90&&b[i]>=65{
			b[i]+=32
		}
	}


	return string(b)
}