package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n > len(flowerbed) {
		return false
	}
	max := 0

	for i := 0; i < len(flowerbed); i++ {
		// 当前格子必须为0
		// 是第一个格子或者前一个格子为0
		// 是最后一个格子或者后一个格子为0
		//这里是因为 i==0和i==len-1不可能同时成立 所以可以合在一起写
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			max++
		}

	}

	return max >= n
}
