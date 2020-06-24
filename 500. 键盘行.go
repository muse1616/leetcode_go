package leetcode

func findWords(words []string) []string {
	hash := map[byte]byte{
		'q': 1, 'Q': 1,
		'w': 1, 'W': 1,
		'e': 1, 'E': 1,
		'r': 1, 'R': 1,
		't': 1, 'T': 1,
		'y': 1, 'Y': 1,
		'u': 1, 'U': 1,
		'i': 1, 'I': 1,
		'o': 1, 'O': 1,
		'p': 1, 'P': 1,
		'a': 2, 'A': 2,
		's': 2, 'S': 2,
		'd': 2, 'D': 2,
		'f': 2, 'F': 2,
		'g': 2, 'G': 2,
		'h': 2, 'H': 2,
		'j': 2, 'J': 2,
		'k': 2, 'K': 2,
		'l': 2, 'L': 2,
		'z': 3, 'Z': 3,
		'x': 3, 'X': 3,
		'c': 3, 'C': 3,
		'v': 3, 'V': 3,
		'b': 3, 'B': 3,
		'n': 3, 'N': 3,
		'm': 3, 'M': 3,
	}
	var res []string
	for _,v:=range words{
		f:=true
		for i:=len(v)-1;i>0;i--{
			if hash[v[i]]!=hash[v[i-1]]{
				f = false
				break
			}
		}
		if f==true{
			res = append(res,v)
		}
	}
	return res

}
