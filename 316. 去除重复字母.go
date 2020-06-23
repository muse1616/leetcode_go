package leetcode

import "bytes"

func removeDuplicateLetters(s string) string {
	b := []byte(s)
	//模拟栈
	var r []byte
	for i, c := range b {
		//栈为空 直接压入
		if len(r) == 0 {
			r = append(r, c)
			continue
		}
		// 当前元素在栈中已经存在 跳过
		if bytes.IndexByte(r, c) != -1 {
			continue
		}
		//栈非空 当前元素小于栈顶元素  栈顶元素在原字符串后还存在
		for len(r) > 0 && c < r[len(r)-1] && bytes.IndexByte(b[i+1:], r[len(r)-1]) != -1 {
			//去掉栈顶元素
			r = r[:len(r)-1]
		}
		// 入栈
		r = append(r, c)
	}

	return string(r)

}
