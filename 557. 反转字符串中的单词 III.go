package leetcode

func reverseWords(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); {
		//跳过空格
		if b[i] == ' ' {
			continue
		}
		//找下个空格位置
		indexSpace := i
		for j := i + 1; j < len(b); j++ {
			if b[j] == ' ' {
				indexSpace = j
				break
			}
			if j == len(b)-1 {
				indexSpace = len(b)
			}
		}
		//翻转i到空格内的字母
		left:=i
		right:=indexSpace-1
		for left<right{
			b[left],b[right] = b[right],b[left]
			left++
			right--
		}
		i = indexSpace + 1
	}

	return string(b)
}
