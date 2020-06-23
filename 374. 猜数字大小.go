package leetcode

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
//题目预定义接口 无视
func guess(num int) int { return 0 }

func guessNumber(n int) int {
	left:=1
	right:=n
	for left<=right{
		mid:=(left+right)/2
		if guess(mid)==0{
			return mid
		}else if guess(mid)==-1{
			right=mid-1
		}else {
			left = mid+1
		}
	}
	return 0
}
