package leetcode

import "strings"

func numUniqueEmails(emails []string) int {
	hash := make(map[string]bool, 10)
	for _, v := range emails {
		// arr[0]本地名称
		// arr[1]域名
		arr := strings.Split(v, "@")
		//	处理本地名称
		// 去掉点 截取到+号前的内容
		index := strings.Index(arr[0], "+")
		if index != -1 {
			arr[0] = arr[0][:index]
		}
		arr[0] = strings.ReplaceAll(arr[0], ".", "")

		f := arr[0] + "@" + arr[1]
		if _, ok := hash[f]; !ok {
			hash[f] = true
		}

	}
	return len(hash)
}
