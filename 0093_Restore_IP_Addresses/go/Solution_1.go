package main

import (
	"fmt"
	"strconv"
)

// 回溯法

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)

	// 检验ip段是否符合，ip段范围[0, 255]且不能有前导0
	var isIp func(ip string) bool
	isIp = func(ip string) bool {
		if len(ip) == 1 || (len(ip) > 1 && ip[0:1] != "0") {
			num, _ := strconv.Atoi(ip)
			if num >= 0 && num <= 255 {
				return true
			}
		}

		return false
	}

	var backtrack func(start, cnt int, s, subStr string)
	backtrack = func(start, cnt int, s, subStr string) {
		if start == len(s) && cnt == 4 {
			subStr = subStr[1:] // 去除最前端的"."
			res = append(res, subStr)
			return
		}

		for i := 1; i <= 3; i++ {
			end := start + i
			if end > len(s) {
				break
			}
			ip := s[start:end]
			// 剪枝 cnt < 4
			if isIp(ip) && cnt < 4 {
				nextSubStr := fmt.Sprintf("%s.%s", subStr, ip)
				backtrack(end, cnt+1, s, nextSubStr)
			}
		}
	}

	backtrack(0, 0, s, "")
	return res
}
