package leetcode

import "math"

func myPow(x float64, n int) float64 {
	if n==0{
		return 1
	}else if n <0{
		return 1 / myPow(x,-n)
	}else if n>0&&n%2==0{
		return myPow(x*x,n/2)
	}else {
		return myPow(x,n-1)*x
	}
}
