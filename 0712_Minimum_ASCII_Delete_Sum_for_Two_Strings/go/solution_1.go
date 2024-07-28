package main

// 思路: 类似求两个字符串的编辑距离 lc0042、lc0583

// 时间复杂度: O(mn)    空间复杂度: O(mn)

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)

	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = dp[0][i-1] + int(s2[i-1])
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// dp[i-1][j]表示删除s1的第i个字符
				// dp[i][j-1]表示删除s2的第j个字符
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}

	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
