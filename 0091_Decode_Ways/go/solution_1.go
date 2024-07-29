package main

// 思路: 动态规划
// 类似lc2466
// 类似爬楼梯问题  dp[i]可以由dp[i-1]演变而来，也可以由dp[i-2]演变而来,通过条件由dp[i-1]和dp[i-2]来获得dp[i]

// 时间复杂度: O(n)    空间复杂度: O(n)

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0]-'0' == 0 {
		return 0
	}

	// dp[i] = k, 表示以i结尾的字符串可以表示的解码总数
	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		// s[i]能够正确解码
		if s[i]-'0' != 0 {
			dp[i] = dp[i-1]
		}

		// s[i-1]和s[i]的组合能够正确解码
		if (s[i-1]-'0')*10+(s[i]-'0') >= 10 && (s[i-1]-'0')*10+(s[i]-'0') <= 26 {
			if i > 1 {
				dp[i] += dp[i-2]

			} else { // 开头的前两位s[0]s[1]
				dp[i] += 1
			}
		}
	}

	return dp[n-1]
}
