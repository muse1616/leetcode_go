package leetcode

import (
	"strconv"
)

type Stack []int

func (stack *Stack) push(value int) {
	*stack = append(*stack, value)
}

func (stack *Stack) pop() (t int) {
	theStack := *stack
	if len(theStack) == 0 {
		return
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value
}

//栈完成逆波兰表达式
func evalRPN(tokens []string) int {
	s := Stack{}
	//数字入栈 符号弹两个进行运算 然后入栈
	for i := 0; i < len(tokens); i++ {
		if tokens[i] != "+" && tokens[i] != "*" && tokens[i] != "-" && tokens[i] != "/" {
			val, _ := strconv.Atoi(tokens[i])
			s.push(val)
		} else {
			if tokens[i] == "+" {
				a1 := s.pop()
				a2 := s.pop()
				s.push(a1 + a2)
			} else if tokens[i] == "*" {
				a1 := s.pop()
				a2 := s.pop()
				s.push(a1 * a2)
			} else if tokens[i] == "/" {
				a1 := s.pop()
				a2 := s.pop()
				s.push(a2 / a1)
			} else if tokens[i] == "-" {
				a1 := s.pop()
				a2 := s.pop()
				s.push(a2 - a1)
			}
		}
	}
	return s.pop()
}
