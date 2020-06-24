package leetcode

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	var rank []string
	//复制一份
	tmp:=make([]int,len(nums))
	copy(tmp,nums)
	//升序排序
	sort.Ints(nums)
	//名次hash
	hash := make(map[int]string, 0)
	//从后向前遍历
	for i := len(nums) - 1; i >= 0; i-- {
		hash[nums[i]] = strconv.Itoa(len(nums)-i)
	}
	for i:=0;i<len(tmp);i++{
		if hash[tmp[i]]=="1"{
			rank = append(rank,"Gold Medal")
		}else if hash[tmp[i]]=="2"{
			rank = append(rank,"Silver Medal")
		}else if hash[tmp[i]]=="3"{
			rank = append(rank,"Bronze Medal")
		}else {
			rank = append(rank,hash[tmp[i]])
		}
	}
	return rank
}
