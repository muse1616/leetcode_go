package leetcode

func checkRecord(s string) bool {
	b := []byte(s)
	count := 0
	lateCount:=0
	for i := 0; i < len(b); i++ {
		if b[i] == 'P' {
			lateCount=0
		}else if b[i] == 'A' {
			if count == 1 {
				return false
			} else {
				count++
				lateCount=0
			}
		} else if b[i]=='L'{
			lateCount++
			if lateCount>=3{
				return false
			}
		}
	}
	return true
}
