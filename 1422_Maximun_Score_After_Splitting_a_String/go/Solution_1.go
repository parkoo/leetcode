package main

import "strings"

// 暴力法
// 时间复杂度：O(n^2)  空间复杂度：O(1)

func maxScore_1(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	}

	for i := 1; i <= len(s)-1; i++ {
		temp := strings.Count(s[:i], "0") + strings.Count(s[i:len(s)], "1")
		if temp > res {
			res = temp
		}
	}

	return res
}
