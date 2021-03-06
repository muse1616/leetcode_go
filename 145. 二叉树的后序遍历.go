package leetcode

import "errors"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 非递归写法
type Stack []*TreeNode

func (stack *Stack) push(value *TreeNode) {
	*stack = append(*stack, value)
}

func (stack *Stack) pop() (t *TreeNode, e error) {
	theStack := *stack
	if len(theStack) == 0 {
		e = errors.New("empty stack")
		return
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

func (stack Stack) Top() (*TreeNode, error) {
	if len(stack) == 0 {
		return nil, errors.New("empty")
	}
	return stack[len(stack)-1], nil
}

// 非递归写法 栈迭代模拟递归即可
func postorderTraversal(root *TreeNode) []int {
	//结果
	var result []int
	if root == nil {
		return result
	}
	//栈
	s := Stack{}
	//root入栈
	s.push(root)
	//栈非空时
	for len(s) != 0 {
		t, _ := s.Top()
		//	弹出
		_, _ = s.pop()
		if t != nil {
			s.push(t)
			//在当前节点之前加入一个空节点表示已经访问过了
			s.push(nil)

			if t.Right != nil {
				s.push(t.Right)
			}
			if t.Left != nil {
				s.push(t.Left)
			}

		} else {
			top, _ := s.Top()
			result = append(result, top.Val)
			_, _ = s.pop()
		}
	}

	return result
}
