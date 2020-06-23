package leetcode

func toHex(num int) string {
	// 如果是负数，转换成补码形式，比如：-1 + 4294967296 = 4294967295 = 0xffffffff
	if num < 0 {
		num += 4294967296
	}
	//答案
	var ans []rune
	//保存十六进制字符
	hash := [...]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	for {
		t := num % 16
		num /= 16
		//注意头插
		ans = append([]rune{hash[t]}, ans...)
		if num == 0 {
			break
		}
	}

	return string(ans)
}
