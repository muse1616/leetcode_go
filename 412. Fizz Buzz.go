package leetcode

import "strconv"

func fizzBuzz(n int) []string {
	var result []string
	for i:=0;i<=n;i++{
		if i>=15&&i%15 == 0{
			result = append(result,"FizzBuzz")
			continue
		}
		if i%5 == 0 && i >=5{
			result = append(result,"Buzz")
			continue
		}
		if i%3 == 0 && i>=3{
			result = append(result,"Fizz")
			continue
		}
		result  = append(result,strconv.Itoa(i))
	}


	return result
}