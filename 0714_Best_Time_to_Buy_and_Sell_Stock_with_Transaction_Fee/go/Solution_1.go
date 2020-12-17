package main

// 动态规划
// 时间复杂度：O(n) 空间复杂度：O(n)

func maxProfit_1(prices []int, fee int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	// 状态初始化
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	// 状态转移 dp[i][1] 表示进行到第i天时，选择持有股票的最大收益，dp[i][0] 表示进行到第i天时，选择不持有股票的最大收益
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}

	return dp[n-1][0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
