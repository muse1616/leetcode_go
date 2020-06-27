package leetcode

import "strings"

type Trie struct {
	isEnd int
	next  [26]*Trie
}

func Constructer() Trie {
	return Trie{}
}
func (this *Trie) Insert(word string, index int) {
	for _, c := range word {
		if this.next[c-'a'] == nil {
			this.next[c-'a'] = new(Trie)
		}
		this = this.next[c-'a']
	}
	this.isEnd = index
}

//搜索最长的单词
func (this *Trie) Search(words []string) string {
	result := ""
	stack := []*Trie{this}
	for len(stack) > 0 {
		//	栈顶元素
		top := stack[len(stack)-1]
		//	出栈
		stack = stack[:len(stack)-1]
		// 栈顶元素节点单词存在或者栈顶元素就是根元素
		if top.isEnd > 0 || top == this {
			if top != this {
				//需要减1 因为top.end从1开始存放
				currentWord := words[top.isEnd-1]
				//当前单词长度大 或长度相同字典序小
				if len(currentWord) > len(result) || (len(currentWord) == len(result) && strings.Compare(result, currentWord) == 1) {
					result = currentWord
				}
			}
			for _, v := range top.next {
				if v != nil {
					stack = append(stack, v)
				}
			}

		}

	}
	return result
}

func longestWord(words []string) string {
	trie := Constructer()
	//加入前缀树
	for index, word := range words {
		trie.Insert(word, index+1)
	}
	return trie.Search(words)
}
