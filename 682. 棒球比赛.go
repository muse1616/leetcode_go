package leetcode

import "strconv"

type Stack []int

// 删除操作
func (s *Stack) Pop() {
	old := *s
	*s = old[:len(old)-1]
}

// 取出操作
func (s Stack) Top() int {
	if len(s) == 0 {
		return 0
	}
	return s[len(s)-1]
}

// 入栈操作
func (s *Stack) Push(a int) {
	old := *s
	*s = append(old, a)
}

// 求长度
func (s Stack) Len() int {
	return len(s)
}

func calPoints(ops []string) int {
	res := Stack{}
	for i := 0; i < len(ops); i++ {
		if ops[i] == "C" {
			res.Pop()
		} else if ops[i] == "D" {
			res.Push(res.Top() * 2)
		} else if ops[i] == "+" {
			temp1 := res.Top()
			res.Pop()
			temp2 := res.Top()
			res.Push(temp1)
			res.Push(temp1 + temp2)
		} else {
			temp, _ := strconv.Atoi(ops[i])
			res.Push(temp)
		}
	}
	result := 0
	for res.Len() != 0 {
		result += res.Top()
		res.Pop()
	}
	return result
}
