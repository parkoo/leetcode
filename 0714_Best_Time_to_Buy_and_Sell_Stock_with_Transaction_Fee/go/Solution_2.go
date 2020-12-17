package main

// 动态规划, 当天状态只与前一天的状态有关，使用两个变量即可维护转移关系
// 时间复杂度：O(n), 空间复杂度：O(1)

func maxProfit_2(prices []int, fee int) int {
	n := len(prices)

	// 状态初始化
	sell := 0
	buy := -prices[0]

	// 状态转移，等号左边表示当天的状态，等号右边表示前一天的状态
	for i := 1; i < n; i++ {
		sell = max_2(sell, buy+prices[i]-fee)
		buy = max_2(buy, sell-prices[i])
	}

	return sell
}

func max_2(i, j int) int {
	if i > j {
		return i
	}
	return j
}
