package leetcode

import "sort"

var graph map[string][]string
var ans []string

func findItinerary(tickets [][]string) []string {
	graph = make(map[string][]string)
	ans = make([]string, 0)
	//建立图邻接表
	for _, v := range tickets {
		graph[v[0]] = append(graph[v[0]], v[1])
	}
	//将邻接表按字典序排序 dfs每次都取出第一个即可
	for _, v := range graph {
		sort.Strings(v)
	}
	dfs("JFK")
	return ans
}
func dfs(airport string) {
	for len(graph[airport]) != 0 {
		//	去除该节点
		v := graph[airport][0]
		graph[airport] = graph[airport][1:]
		dfs(v)
	}
	//因为先找不到路径的一定是在答案的末尾 这样后添加的就在先添加的前面 所以此处是放在前面
	ans = append([]string{airport}, ans...)
}
