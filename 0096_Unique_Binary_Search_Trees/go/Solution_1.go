package main

// 动态规划：dp[i] = dp[j]*dp[i-1-j]
// 时间复杂度：O(n^2)  空间复杂度：O(n)

func numTrees_1(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}

	return dp[n]
}
