package main

// 思路: LCS问题（最长公共子序列）经典二维动态规划问题
// LCS(m,n)表示s1[0...m]和s2[0...n]的最长公共子序列的长度
// 若s1[m] == s2[n]：
// LCS(m,n) = 1 + LCS(m-1,n-1)
// 若s1[m] != s2[n]：
// LCS(m,n) = max( LCS(m-1,n), LCS(m, n-1) ) (s1往前缩一缩或者s2往前缩一缩)

// 时间复杂度: O(m*n)    空间复杂度: O(m*n)

func longestCommonSubsequence_1(text1 string, text2 string) int {
	s1, s2 := []rune(text1), []rune(text2)
	m, n := len(s1), len(s2)

	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
