package leetcode

import (
	"strconv"
	"strings"
)

var hash map[string]int

func subdomainVisits(cpdomains []string) []string {
	var result []string
	hash = make(map[string]int, 10)
	for _, v := range cpdomains {
		str := strings.Split(v, " ")
		n, _ := strconv.Atoi(str[0])
		count(n, str[1])
	}
	//hash -> []string
	for k, v := range hash {
		str := strconv.Itoa(v) + " " + k
		result = append(result, str)
	}
	return result
}

//次数 域名 补齐hash表
func count(num int, domain string) {
	//自身访问
	hash[domain] += num
	//	几个点号
	if strings.Index(domain, ".") == strings.LastIndex(domain, ".") {
		//	只出现一次
		index := strings.Index(domain, ".")
		hash[domain[index+1:]] += num
	} else {
		//	出现两次
		index := strings.Index(domain, ".")
		hash[domain[index+1:]] += num
		lastIndex := strings.LastIndex(domain, ".")
		hash[domain[lastIndex+1:]] += num
	}

}
