package leetcode

func longestPalindrome(s string) string {
	var output []byte
	bs := []byte(s)
	if len(bs) > 0 {
		output = bs[0:1]
	}
	for i := 0; i < len(bs); i++ {
		for j := i + len(output); j < len(bs); j++ {
			flag := 0
			for k := 0; k < (j-i+1)/2; k++ {
				if bs[k+i] != bs[j-k] {
					flag = 1
					break
				}
			}
			if flag == 0 && j-i+1 > len(output) {
				output = bs[i : j+1]
			}
		}
	}
	return string(output)
}
