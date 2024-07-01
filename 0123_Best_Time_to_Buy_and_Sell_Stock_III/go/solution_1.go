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

	// 状态依赖: sell1 -> buy2 -> sell2
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i]) // 这里的buy1, 直接依赖上一步算出的结果
		buy2 = max(buy2, sell1-prices[i])  // 这里的sell1, 直接依赖上一步算出的结果
		sell2 = max(sell2, buy2+prices[i]) // 这里的buy2, 直接依赖上一步算出的结果
	}

	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
