package main

// 动态规划 二维数组
// 时间复杂度：O(n)  空间复杂度：O(n)

func maxProfit_1(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) // 可以多次交易与只许单次交易的区别
	}

	return dp[n-1][0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
