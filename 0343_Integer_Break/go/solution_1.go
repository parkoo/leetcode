package main

// 思路: 动态规划

// 时间复杂度: O(n^2)    空间复杂度: O(n)

func integerBreak(n int) int {

	// dp[i] = k, 表示将i拆分成至少2个正整数后,这些正整数的最大积为k
	// 将i拆分成两部分j和(i-j)
	// 为了保证至少拆分成2个正整数，可以分为两种情况
	// 1. j不拆分且(i-j)也不拆分，此时 dp[i] = j*(i-j)
	// 2. j不拆分但(i-j)作拆分， 此时 dp[i] = j*dp[i-j] (dp[i-j]表示把(i-j)至少拆分成两部分的最大积)
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(max(dp[i], j*dp[i-j]), j*(i-j))
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
