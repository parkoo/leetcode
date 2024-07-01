package main

// 优化空间的一维动态规划
// 时间复杂度：O(n)  空间复杂度：O(1)

func maxProfit_3(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	// dp[0]表示当天手中不持股票时可获的最大利润
	// dp[1]表示天手中持有股票时可获的最大利润
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = -prices[0]

	// ”当天不持有”可以由“前天不持有”以及“前天持有当天卖出“两种情况获得
	// ”当天持有”可以由“前天持有”以及“前天不持有当天买入“两种情况获得
	for i := 1; i < n; i++ {
		// 等号右边的dp[0]表示前一天天不持有，等号左边的dp[0]表示当天不持有，由等号右边的dp[i]获得
		dp[0] = max(dp[0], dp[1]+prices[i])
		dp[1] = max(dp[1], -prices[i]) // 当天买入后，当天最大利润为 -prices[i]
	}

	return dp[0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
