package main

// 动态规划  dp[i]表示登上第i阶台阶的方法总数
// 时间复杂度：O(n)  空间复杂度：O(n)

func climbStairs_1(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i:=2; i<=n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}