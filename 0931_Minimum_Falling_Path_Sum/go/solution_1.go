package main

// 思路: 二维区域 动态规划

// 时间复杂度: O(n^2)    空间复杂度: O(n^2)

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[0][i] = matrix[0][i]
	}

	for i := 1; i < n; i++ {
		dp[i][0] = min(dp[i-1][0], dp[i-1][1]) + matrix[i][0]
		dp[i][n-1] = min(dp[i-1][n-1], dp[i-1][n-2]) + matrix[i][n-1]
		for j := 1; j < n-1; j++ {
			dp[i][j] = min(dp[i-1][j+1], min(dp[i-1][j-1], dp[i-1][j])) + matrix[i][j]
		}
	}

	minVal := dp[n-1][0]
	for i := 0; i < n; i++ {
		minVal = min(minVal, dp[n-1][i])
	}

	return minVal
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
