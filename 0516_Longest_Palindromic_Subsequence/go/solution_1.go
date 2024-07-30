package main

// 思路: 子序列问题  全局性考虑

// 注意区别lc0005, 其为寻找最长回文子串（中间不能删去字符，需要连续）, 连续的子串还可以用中心扩展法
// 而本题寻找最长回文子序列（中间可以删去一些字符）

// 时间复杂度: O(n^2)    空间复杂度: O(n^2)

func longestPalindromeSubseq_1(s string) int {
	n := len(s)

	// dp[i][j]=k, 表示 s[i:j+1] 区间内的最长回文子序列的长度为k（s[i]、s[j] 可以是回文子序列的一部分也可以不是）
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for end := 0; end < n; end++ {
		dp[end][end] = 1
		for start := end - 1; start >= 0; start-- {
			if s[start] == s[end] {
				dp[start][end] = dp[start+1][end-1] + 2

			} else {
				dp[start][end] = max(dp[start+1][end], dp[start][end-1])
			}
		}
	}

	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
