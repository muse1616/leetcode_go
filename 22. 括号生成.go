package leetcode
func generateParenthesis(n int) []string {
	result:=make([]string, 0,10)
	add("", n, n,&result)
	return result



}
// 回溯
func add(answer string,left int,right int,result *[]string){
	// 剩下的左括号大于有括号 即已经用的右括号小于左括号
	if left >right{
		return
	}
	if left==0&&right==0{
		*result = append(*result,answer)
		return
	}

	// 优先左
	if left>0{
		add(answer+"(",left-1,right,result)
	}
	if right>0{
		add(answer+")",left,right-1,result)
	}
}