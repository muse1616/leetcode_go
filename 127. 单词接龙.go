package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// word hash 快速查找删除
	dict := make(map[string]bool)
	for _, word := range wordList {
		dict[word] = true
	}
	// hash中不存在目标单词 直接返回0
	if _, ok := dict[endWord]; !ok {
		return 0
	}
	q1 := make(map[string]bool)
	q2 := make(map[string]bool)
	q1[beginWord] = true // 头
	q2[endWord] = true   // 尾
	//由题 每个单词长度都相同
	l := len(beginWord)
	//步长
	steps := 0
	//双向BFS 两个队列都非空
	for len(q1) > 0 && len(q2) > 0 {
		steps++
		// 永远都将q1保持为较短的队列 表示q2中的单词都遍历过 只要q1再有q2中出现过的替换一个字母的单词 就找到了
		if len(q1) > len(q2) {
			q1, q2 = q2, q1
		}
		//临时队列
		q := make(map[string]bool)
		//遍历 q1
		for k := range q1 {
			//当前单词
			chs := []rune(k)
			for i := 0; i < l; i++ {
				ch := chs[i]
				// 对每一位从a-z尝试
				for c := 'a'; c <= 'z'; c++ {
					// 替换字母组成新的单词
					chs[i] = c
					t := string(chs)
					// 看新单词是否在q2中 如果在的话说明已经找到了
					if _, ok := q2[t]; ok {
						return steps + 1
					}
					// 看新单词是否在hash中
					if _, ok := dict[t]; !ok {
						// 不在hash就跳出循环
						continue
					}
					//若在字典中就删除 表示访问过了
					delete(dict, t)
					//加入临时队列
					q[t] = true
				}
				//还原新单词
				chs[i] = ch
			}
		}
		// q1 更新为 临时队列
		q1 = q
	}
	// 未找到
	return 0
}
