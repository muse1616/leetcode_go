package leetcode

import "fmt"

func readBinaryWatch(num int) []string {
	var result []string
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if bitNums(i)+bitNums(j) == num {
				result = append(result, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return result
}

//计算当前时间会亮几盏灯 即二进制有几个零
func bitNums(i int) int {
	count := 0
	for i > 0 {
		// 末尾有1即模2余1
		if i%2 == 1 {
			count++
		}
		//舍去低位
		i = i >> 1
	}
	return count
}
