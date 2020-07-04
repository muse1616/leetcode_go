package leetcode

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	//当余数出现循环的时候，对应的商也会循环
	if numerator == 0 || denominator == 0 {
		return "0"
	}
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}
	var ans []byte
	//判断符号
	if (numerator > 0 && denominator < 0) || (numerator < 0 && denominator > 0) {
		ans = append(ans, '-')
	}
	//都转换成正数计算
	numerator = abs(numerator)
	denominator = abs(denominator)
	// 整数部分
	ans = append(ans, []byte(strconv.Itoa(numerator/denominator))...)
	numerator = numerator % denominator
	ans = append(ans, '.')
	//小数部分
	// 记录数字出现的位置 出现过就说明循环了
	hash := make(map[int]int)
	// 记录循环开始的位置
	startPos := -1
	// 小数部分
	var low []byte
	for numerator > 0 {
		// 出现循环
		if pos, ok := hash[numerator]; ok {
			startPos = pos
			break
		}
		// 记录位置
		hash[numerator] = len(low)
		// 分母需要乘以10 因为得到的数字需要是0-9的数字
		numerator = numerator * 10
		// 加入小数部分
		low = append(low, []byte(strconv.Itoa(numerator/denominator))...)
		// 取余
		numerator = numerator % denominator
	}
	//加括号
	if startPos != -1 {
		// 添加循环出现前的的部分 加左括号 之后部分 右括号
		ans = append(ans, low[:startPos]...)
		ans = append(ans, '(')
		ans = append(ans, low[startPos:]...)
		ans = append(ans, ')')
	} else {
		// 没有循环出现 直接添加
		ans = append(ans, low...)
	}

	return string(ans)

}
func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}
