package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	//var min byte
	//min = 128
	//aim := letters[0]
	//for _, v := range letters {
	//	if v-target > 0 {
	//		min = byte(math.Min(float64(min), float64(v-target)))
	//	}
	//	if min == v-target{
	//		aim = v
	//	}
	//}
	//return aim

	left, right := 0, len(letters)-1
	if letters[right] <= target {
		return letters[left]
	}
	for left < right {
		middle := left + (right-left)/2
		if letters[middle] > target {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return letters[left]

}
