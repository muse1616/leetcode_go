package leetcode

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	var answer []int
	//求出初始偶数值的和
	sum:=0
	for _,v:=range A{
		if v % 2==0{
			sum+=v
		}
	}
	//原来是奇数 加完后是奇数 偶数值不变  原来是奇数 加完后是偶数 偶数值加这个偶数
	//原来是偶数 加完后是偶数 加这个加的值 原来是偶数 加完后是奇数 偶数值减原来的偶数
	for _,v:=range queries{
		if A[v[1]] % 2 !=0 && (A[v[1]]+v[0]) %2!=0{
			A[v[1]]+=v[0]
			answer = append(answer,sum)
		}else if A[v[1]]%2!=0 && (A[v[1]]+v[0])%2==0{
			A[v[1]]+=v[0]
			sum+=A[v[1]]
			answer = append(answer,sum)
		}else if A[v[1]]%2==0 && (A[v[1]]+v[0])%2==0{
			A[v[1]]+=v[0]
			sum+=v[0]
			answer = append(answer,sum)
		}else if A[v[1]]%2==0 && (A[v[1]]+v[0])%2!=0{
			sum-=A[v[1]]
			A[v[1]]+=v[0]
			answer = append(answer,sum)

		}
	}


	return answer
}
