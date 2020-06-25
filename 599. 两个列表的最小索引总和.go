package leetcode

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	hash1 := make(map[string]int)
	hash2 := make(map[string]int)
	hash := make(map[string]int)
	for k, v := range list1 {
		hash1[v] = k
	}
	for k, v := range list2 {
		hash2[v] += k
	}
	for k, v1 := range hash1 {
		if v2, ok := hash2[k]; ok {
			hash[k] = v1 + v2
		}
	}
	min := 3000
	//	找到最小索引
	for _, v := range hash {
		min = int(math.Min(float64(v), float64(min)))
	}
	var result []string
	for k, v := range hash {
		if v == min {
			result = append(result, k)
		}

	}
	return result
}
