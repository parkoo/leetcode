package main

// 二维动态规划
// 时间复杂度：O(n)  空间复杂度：O(n)

func maxProfit_2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	// dp[i][0]表示第i天手中不持股票时可获的最大利润
	// dp[i][1]表示第i天手中持有股票时可获的最大利润
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	// ”当天不持有”可以由“前天不持有”以及“前天持有当天卖出“两种情况获得
	// ”当天持有”可以由“前天持有”以及“前天不持有当天买入“两种情况获得
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i]) // 当天买入后，当天最大利润为 -prices[i]
	}

	return dp[n-1][0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
