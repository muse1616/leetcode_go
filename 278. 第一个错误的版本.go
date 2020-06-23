package leetcode
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good

 */

//题目给的接口 无视
func isBadVersion(n int)bool  {
	return false
}

func firstBadVersion(n int) int {
	left,right:=1,n
	for left<right{
		mid:=(right+left)/2
		if isBadVersion(mid){
			right = mid
		}else {
			left = mid+1
		}
	}
	return left
}