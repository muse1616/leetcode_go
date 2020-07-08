package leetcode

func isValidSerialization(preorder string) bool {
	slot := 1
	for i := 0; i < len(preorder); i++ {
		if preorder[i] == ',' {
			//每个节点消耗一个槽位
			slot--
			if slot < 0 {
				return false
			}
			//	非空增加两个
			if preorder[i-1] != '#' {
				slot += 2
			}

		}
	}
	//最后一个节点单独处理
	if preorder[len(preorder)-1]!='#'{
		slot++
	}else {
		slot--
	}
	//最后一定用完
	return slot == 0
}
