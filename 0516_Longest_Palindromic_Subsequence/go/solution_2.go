package main

// 思路: 子序列问题
// "最长回文子序列问题" 可以转换为"1143两个序列的最长公共子序列问题"
// 若将s进行翻转得到t, 则s的最长回文子序列 == lcs(s,t)

// 注意区别lc0005, 其为寻找最长回文子串（中间不能删去字符，需要连续）, 连续的子串还可以用中心扩展法
// 而本题寻找最长回文子序列（中间可以删去一些字符）

// 时间复杂度: O(n^2)    空间复杂度: O(n^2)

func longestPalindromeSubseq_2(s string) int {
	n := len(s)

	ts := make([]rune, n)
	for i, c := range []rune(s) {
		ts[n-1-i] = c
	}
	t := string(ts)

	return lcs(s, t)
}

// 求s1和s2的最长公共子序列
func lcs(s1, s2 string) int {
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
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
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
