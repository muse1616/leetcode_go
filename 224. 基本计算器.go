package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

//先转成逆波兰 再用逆波兰算法

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	//将中缀转为逆波兰 然后按照逆波兰表达式求值即可
	return evalRPN(toReversePolish(s))
}

func toReversePolish(s string) (tokens []string) {

	stack := Stack2{}
	tmp := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if tmp != "" {
				tokens = append(tokens, tmp)
				tmp = ""
			}

			stack.push(string(s[i]))
		} else if s[i] == '+' || s[i] == '-' {
			if tmp != "" {
				tokens = append(tokens, tmp)
				tmp = ""
			}
			//	弹出之前的+-
			for ; len(stack) > 0 && (stack.top() == "+" || stack.top() == "-"); {
				tokens = append(tokens, stack.pop())
			}
			stack.push(string(s[i]))
		} else if s[i] == ')' || i == len(s)-1 {

			if i == len(s)-1 && s[i] != ')' {
				tmp += string(s[i])
				if tmp != "" {
					tokens = append(tokens, tmp)
					tmp = ""
				}
			} else {
				if tmp != "" {
					tokens = append(tokens, tmp)
					tmp = ""
				}
			}

			for {
				if len(stack) <= 0 {
					break
				}
				a := stack.pop()
				if a == "(" {
					break
				} else {
					tokens = append(tokens, a)
				}
			}
		} else {
			tmp += string(s[i])
		}
	}
	for len(stack) != 0 {
		tokens = append(tokens, stack.pop())
	}
	return tokens
}

//栈完成逆波兰表达式
func evalRPN(tokens []string) int {
	fmt.Println(tokens)
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

type Stack2 []string

func (stack *Stack2) push(value string) {
	*stack = append(*stack, value)
}
func (stack Stack2) top() string {
	return stack[len(stack)-1]
}
func (stack *Stack2) pop() (t string) {
	theStack := *stack
	if len(theStack) == 0 {
		return
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value
}
