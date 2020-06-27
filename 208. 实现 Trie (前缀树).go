package leetcode

type Trie struct {
	//该节点是否是一个串的结束
	isEnd bool
	//字母映射表
	next [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, c := range word {
		//开辟新节点
		if this.next[c-'a'] == nil {
			this.next[c-'a'] = new(Trie)
		}
		this = this.next[c-'a']
	}
	//结尾符号
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, c := range word {
		//没有节点
		if this.next[c-'a'] == nil {
			return false
		}
		this = this.next[c-'a']
	}
	if this.isEnd == false {
		//	非结束标志
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		if this.next[c-'a'] == nil {
			return false
		}
		this = this.next[c-'a']
	}
	//判断前缀时无需结束标志
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
