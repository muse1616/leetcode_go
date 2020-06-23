package leetcode

import "strings"

func simplifyPath(path string) string {
	// 按照斜杠分割
	buf := strings.Split(path, "/")
	// 模拟栈
	var stack []string
	for i := 0; i < len(buf); i++ {
		if buf[i] == "." || buf[i] == "" {
			continue
		}
		if buf[i] == ".." {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack = append(stack, buf[i])
		}
	}

	return "/" + strings.Join(stack, "/")

}
