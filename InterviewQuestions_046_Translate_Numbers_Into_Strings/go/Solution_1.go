package main

import (
	"strconv"
)

// 动态规划，dp[i] 表示以第i个数字结尾时，可转化为字符串的总数
// 时间复杂度：O(n)  空间复杂度：O(n), 其中 n=log(num)
// 注：给定的num是整数，把它转换成字符串，字符串的长度就是log(num)，以10为底，num的对数。

func translateNum_1(num int) int {
	numStr := strconv.Itoa(num)

	dp := make([]int, len(numStr)+1)
	dp[0] = 1
    dp[1] = 1

	for i:=2; i<=len(numStr); i++ {
		dp[i] = dp[i-1]
			
		if temp, _ := strconv.Atoi(numStr[i-2:i]); temp<26 && temp>9 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(numStr)]
}