package leetcode

func powerfulIntegers(x int, y int, bound int) []int {
	var ans []int
	m := make(map[int]bool)
	powerX, powerY := 1, 1

	for powerX < bound {
		for powerY < bound {
			t := powerX + powerY
			if t <= bound {
				m[t] = true
			}
			powerY *= y
			if powerY == 1 {
				break
			}
		}

		powerX *= x
		powerY = 1
		if powerX == 1 {
			break
		}
	}

	for k := range m {
		ans = append(ans, k)
	}
	return ans
}
