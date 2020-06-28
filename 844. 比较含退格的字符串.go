package leetcode

func backspaceCompare(S string, T string) bool {
	stack1 := make([]byte, 0, 10)
	stack2 := make([]byte, 0, 10)
	for _, v := range S {
		if v == '#' {
			if len(stack1) == 0 {
				continue
			} else {
				stack1 = stack1[:len(stack1)-1]
			}
		} else{
			stack1 = append(stack1,byte(v))
		}
	}
	for _, v := range T {
		if v == '#' {
			if len(stack2) == 0 {
				continue
			} else {
				stack2 = stack2[:len(stack2)-1]
			}
		} else{
			stack2 = append(stack2,byte(v))
		}
	}
	return string(stack1) == string(stack2)
}
