package leetcode

type WordDictionary struct {
	next  [26]*WordDictionary
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	for _, c := range word {
		if this.next[c-'a'] == nil {
			this.next[c-'a'] = new(WordDictionary)
		}
		this = this.next[c-'a']
	}
	//结束符号
	this.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	for k, c := range word {
		if c == '.' {
			//	回溯
			for _, v := range this.next {
				if v != nil && v.Search(word[k+1:]) {
					return true
				}
			}
			// 所有的节点回溯都不符合表达式
			return false
		} else if this.next[c-'a'] == nil {
			// 节点空
			return false
		} else {
			this = this.next[c-'a']
		}

	}
	// 判断结束符号
	return this.isEnd
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
