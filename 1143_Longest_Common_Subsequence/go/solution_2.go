package main

// 思路: LCS问题（最长公共子序列）经典二维动态规划问题, [优化为一维空间]
// LCS(m,n)表示s1[0...m]和s2[0...n]的最长公共子序列的长度
// 若s1[m] == s2[n]：
// LCS(m,n) = 1 + LCS(m-1,n-1)
// 若s1[m] != s2[n]：
// LCS(m,n) = max( LCS(m-1,n), LCS(m, n-1) ) (s1往前缩一缩或者s2往前缩一缩)

// 时间复杂度: O(m)    空间复杂度: O(m*n)

func longestCommonSubsequence_2(text1 string, text2 string) int {
	s1, s2 := []rune(text1), []rune(text2)
	m, n := len(s1), len(s2)

	dp := make([]int, n+1)
	for i := 1; i <= m; i++ {
		pre := 0 // 定义上一轮的初始值，表示dp[i-1][j-1], 初始化为0

		// dp[j]未变化之前永远代表dp[i-1][j]
		for j := 1; j <= n; j++ {
			temp := dp[j] // 记录本轮dp[i-1][j] 在下一轮j循环中为 dp[i-1][j-1]
			if s1[i-1] != s2[j-1] {
				// i先于j操作，dp[j-1]实际表示dp[i][j-1]
				// 左边的dp[j]是j操作前的状态，表示dp[i-1][j]
				// 右边的dp[j]是j操作后的状态，表示dp[i][j]

				// [x ... - o ... x]
				// [x ... + * ^.. x]
				// 如: 若本轮待处理为*， 则 dp[j]: o,  dp[j-1]: +,
				// 对应下一轮j循环待处理为^, 此时记录的 dp[j],即o, 为dp[i-1][j-1]

				dp[j] = max(dp[j-1], dp[j])

			} else {
				dp[j] = pre + 1
			}

			pre = temp // 本轮的dp[i-1][j], 在下一轮j循环中就是dp[i-1][j-1],即下一轮j循环中的pre
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
