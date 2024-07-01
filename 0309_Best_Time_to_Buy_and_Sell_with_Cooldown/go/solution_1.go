package main

// 思路: 动态规划，转态转移

// 时间复杂度: O(n)    空间复杂度: O(1)

func maxProfit(prices []int) int {
	// buy: 当天结束时，手上持有股票的最大收益
	// freeze: 当天结束时，手上不持有股票，且处于冷冻期的最大收益
	// sell: 当天结束时，手上不持有股票，且不处于冷冻期的最大收益
	buy, freeze, sell := -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		// [前一天结束时，手上持有股票] or [前一天结束时，手上不持有股票，且不处于冷冻期] -> [当天结束时，手上持有股票]
		curBuy := max(buy, sell-prices[i])
		// [前一天结束时，手上持有股票]-> [手上不持有股票，且处于冷冻期]
		curFreeze := buy + prices[i]
		// [前一天结束时，手上不持有股票，且不处于冷冻期] or [前一天结束时，手上不持有股票，且处于冷冻期] -> [当天结束时，手上不持有股票，且不处于冷冻期]
		curSell := max(sell, freeze)

		buy, freeze, sell = curBuy, curFreeze, curSell
	}

	return max(sell, freeze)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
