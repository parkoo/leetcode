package main

// 降维动态规划
// 时间复杂度：O(n^2)  空间复杂度：O(n)

func minimumTotal_2(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)

	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		// dp[j]依赖dp[j-1], 故需从后向前更新
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
		}
		dp[0] = dp[0] + triangle[i][0]
	}

	res := dp[0]
	for i := 0; i < n; i++ {
		if dp[i] < res {
			res = dp[i]
		}
	}

	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
