package leetcode

var digitToLetter map[byte][]string = map[byte][]string{
	'2':{"a","b","c"},
	'3':{"d","e","f"},
	'4':{"g","h","i"},
	'5':{"j","k","l"},
	'6':{"m","n","o"},
	'7':{"p","q","r","s"},
	'8':{"t","u","v"},
	'9':{"w","x","y","z"},
}

func letterCombinations(digits string) []string {
	res:=[]string{}
	// 遍历循环数字
	if digits==""{
		return res
	}

	res = append(res,"")

	for i:=0;i<len(digits);i++{
		// 字母
		letters := digitToLetter[digits[i]]
		// 保存res
		tmp:=res
		// 清空res
		res=[]string{}
		// 遍历letter
		for _,letter:=range letters{
			for _,v:=range tmp{
				res =append(res, v+letter)
			}
		}


	}



	return res

}