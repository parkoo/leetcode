package main

// 思路：动态规划, 当天状态只与前一天的状态有关，使用两个变量即可维护转移关系

// 时间复杂度：O(n), 空间复杂度：O(1)

func maxProfit(prices []int, fee int) int {
	// dp0: 表示当天结束时，不持有股票可获得的最大利润
	// dp1: 表示当天结束时，持有股票可获得的最大利润
	// 在每次卖出时提交手续费
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp0 = max(dp0, dp1+prices[i]-fee)
		dp1 = max(dp1, dp0-prices[i]) // 由于没有冷冻期， 这里依赖的dp0直接去上步计算得到的值，可以理解为当天可以买出在买入
	}

	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
