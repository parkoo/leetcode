package main

// 思路:  动态规划

// 时间复杂度: O(mn)    空间复杂度: O(mn)

func numDistinct_1(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}

	// 定义 dp[i][j] = k 表示s[0:i]中包含t[0:j]的个数

	// 若 s[i] == t[j], 则 dp[i][j] 分为两部分
	// 1. s[0:i-1]已经包含了a个t[0:j], 则增加一个s[i]也一定会包含a个t[0:j]
	// 2. s[0:i-1]已经包含了b个t[0:j-1], 由于s[i] == t[j], 故s[0:i]中包含了b个t[0:j]
	// 即：dp[i][j] = dp[i-1][j] + dp[i-1][j-1]

	// 若 s[i] != t[j]
	// 则 s[0:i]中包含t[0:j]的个数与s[0:i-1]中包含t[0:j]的个数相等
	// 即： dp[i][j] = dp[i-1][j]

	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	// t为空字符串时，s的任一子串(包括空串)都包含t
	for i := 0; i <= m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[m][n]
}
