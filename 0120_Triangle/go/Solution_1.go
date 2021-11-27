package main

// 动态规划
// 时间复杂度：O(n^2)  空间复杂度：O(^2)

func minimumTotal_1(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]

	for i := 1; i < m; i++ {
		// 第i行的第一个点只能由第i-1行的第一个点获取
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		k := len(triangle[i])
		for j := 1; j < k-1; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		// 第i行的最后一个点只能由上一行的最后一个点获取
		dp[i][k-1] = dp[i-1][k-2] + triangle[i][k-1]
	}

	minValue := dp[m-1][0]
	for i := 0; i < len(dp[m-1]); i++ {
		if dp[m-1][i] < minValue {
			minValue = dp[m-1][i]
		}
	}

	return minValue
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
