package leetcode

import (
	"strconv"
)

func calculate(s string) int {
	//字符串转为字符串数组
	var arr []string
	tmp := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if tmp != "" {
				arr = append(arr, tmp)
			}
			tmp = ""
			continue
		} else if s[i] == '*' || s[i] == '+' || s[i] == '/' || s[i] == '-' {
			if tmp != "" {
				arr = append(arr, tmp)
			}
			arr = append(arr, string(s[i]))
			tmp = ""
		} else {
			tmp = tmp + string(s[i])
		}
	}
	if tmp != "" {
		arr = append(arr, tmp)
	}
	//	先转化为逆波兰表达式
	tokens := convert(arr)
	// 计算逆波兰表达式
	return evalRPN(tokens)
}

//转换逆波兰
func convert(s []string) []string {
	var tokens []string
	s2 := Stack2{}
	for i := 0; i < len(s); i++ {
		//	如果是左括号直接压入栈
		if s[i] == "(" {
			s2.push(string(s[i]))
		} else if s[i] == ")" {
			//	右括号的话将栈中直到左括号的所有符号全部压出到表达式中
			for s2.top() != "(" {
				tokens = append(tokens, s2.pop())
			}
			// 左括号删除
			s2.pop()
		} else if s[i] == "+" || s[i] == "-" {
			//	符号的话弹出之前大于等于它的运算符
			for len(s2) != 0 && (s2.top() == "+" || s2.top() == "-" || s2.top() == "*" || s2.top() == "/") {
				tokens = append(tokens, s2.pop())
			}
			//	压入s[i]
			s2.push(s[i])

		} else if s[i] == "*" || s[i] == "/" {
			for len(s2) != 0 && (s2.top() == "*" || s2.top() == "/") {
				tokens = append(tokens, s2.pop())
			}
			//	压入s[i]
			s2.push(s[i])
		} else {
			//	数字直接压入表达式
			tokens = append(tokens, s[i])
		}
	}
	// 剩下的压入栈中
	for len(s2) != 0 {
		tokens = append(tokens, s2.pop())
	}
	return tokens
}

//计算逆波兰
func evalRPN(tokens []string) int {
	//模拟栈结构
	s := Stack{}
	for i := 0; i < len(tokens); i++ {
		//	数字直接入栈
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

// 栈数据结构
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
