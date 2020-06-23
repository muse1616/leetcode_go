package leetcode

func convert(s string, numRows int) string {
	// 处理字符串
	strArr := []rune(s)
	if len(strArr) == 1 || numRows == 1 {
		return s
	}
	// 确定整除的数字n
	core := 2 * (numRows - 1)
	var result []rune
	// 需要一个二维数组切片
	arrayName := make([][]rune, numRows)
	for i := 0; i < len(strArr); i++ {
		for j := 0; j < numRows; j++ {
			if (i+j)%core == 0 || (i-j)%core == 0 {
				arrayName[j] = append(arrayName[j], strArr[i])
				// fmt.Println(string(strArr[i]))
			}
		}
	}
	// 准备数组
	// 遍历字符串
	for i := 0; i < len(arrayName); i++ {
		for j := 0; j < len(arrayName[i]); j++ {
			result = append(result, arrayName[i][j])
		}
	}
	return string(result)
}
