package main

// 思路：二维动态规划，dp空间同时也是存储结果空间

// 时间复杂度：O(n^2)  空间复杂度：O(1) 不包含结果存储

func generate_1(numRows int) [][]int {
	dp := make([][]int, numRows, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1, i+1)
	}

	dp[0][0] = 1
	for i := 1; i < numRows; i++ {
		dp[i][0] = 1
		dp[i][i] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}

	return dp
}
