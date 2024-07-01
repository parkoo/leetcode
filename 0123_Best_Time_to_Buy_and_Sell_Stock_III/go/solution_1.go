package main

// 思路: 动态规划， 状态转移方程

// 时间复杂度: O(n)    空间复杂度: O(1)

func maxProfit(prices []int) int {
	// buy1: 当天结束时，状态为“第一次买入”时可以获得的最大利润
	// sell1: 当天结束时，状态为“第一次卖出”时可以获得的最大利润
	// buy2: 当天结束时，状态为“第二次买入”时可以获得的最大利润
	// sell2: 当天结束时，状态为“第二次卖出”时可以获得的最大利润
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0

	for i := 1; i < len(prices); i++ {
		preBuy1, preSell1 := buy1, sell1
		preBuy2, preSell2 := buy2, sell2

		buy1 = max(preBuy1, -prices[i])
		sell1 = max(preSell1, preBuy1+prices[i])
		buy2 = max(preBuy2, preSell1-prices[i])
		sell2 = max(preSell2, preBuy2+prices[i])
	}

	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
