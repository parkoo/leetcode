package main

// 思路: 动态规划，经典二维DP问题
// TODO 尝试结合lc1143/solution_2, 优化空间复杂度为一维空间

// 时间复杂度: O(m*n)    空间复杂度: O(m*n)

func minDistance_1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// dp[i][j] 表示word1[0:i-1] 与 word2[0:j-1] 的最小编辑距离
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化边界
	// dp[i][0] 可以看作word1[0:i-1] 与 “” 的最小编辑距离, 其值为word1[0:i-1]的长度
	// dp[0][i] 可以看作word2[0:i-1] 与 “” 的最小编辑距离, 其值为word2[0:i-1]的长度
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // 不需要编辑

			} else {
				// 对于 待计算的dp[i][j], 考虑以下三种操作
				// dp[i-1][j] : 删, 增加word1, 即删除word2
				// dp[i][j-1] : 增, 增加word2, 即删除word1
				// dp[i-1][j-1] : 改, 修改word1或者word2
				dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i][j-1])) + 1
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
