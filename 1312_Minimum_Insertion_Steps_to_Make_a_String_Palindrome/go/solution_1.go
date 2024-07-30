package main

// 思路: 可以通过最长回文子序列间接求解
// 先求出“字符串中已有的最大回文子序列”, 见lc0516
// 在此基础上，其它的字符都可以通过在对侧添加一个字符来达到回文
// 因此求出剩余的字符数即可
// 如 leetcode 的最长回文子序列为 eee -> (l)ee(tcod)e -> (l)e[doct]e(tcod)e[l]

// 时间复杂度: O(n^2)    空间复杂度: O(n^2)

func minInsertions(s string) int {
	return len(s) - longestPalindromeSubseq(s)
}

// 最长回文子序列
func longestPalindromeSubseq(s string) int {
	n := len(s)

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
