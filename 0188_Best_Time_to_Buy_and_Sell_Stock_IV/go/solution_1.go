package main

import "math"

// 思路: 动态规划  状态转移

// 时间复杂度: O(n*k)    空间复杂度: O(k)

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 || n == 1 {
		return 0
	}
	k = min(k, n/2) // 避免k过大

	// buys[i]  表示处于第i次买入股票的状态
	// sells[i] 表示处于第i次卖出股票的状态
	// 状态依赖关系： sells[i-1] -> buys[i] -> sells[i]
	buys, sells := make([]int, k+1), make([]int, k+1)

	// 初始化状态: buys[...] = math.MinInt  sells[...] = 0
	for i := 0; i <= k; i++ {
		buys[i] = math.MinInt
	}

	for i := 0; i < len(prices); i++ { // 计算第i天的状态
		for j := 1; j <= k; j++ {
			// 当天处于‘购买第i次股票’，可以由前一天就处于‘购买第i次股票’ 或者 当天处于‘卖出第i-1次’股票两种状态获得
			buys[j] = max(buys[j], sells[j-1]-prices[i])

			// 当天处于‘卖出第i次股票’，可以由前一天就处于‘卖出第i次股票’ 或者 当天处于‘买入第i次’股票两种状态获得
			// 注意这里的buys[j] 取得就是上一行刚算出的值
			sells[j] = max(sells[j], buys[j]+prices[i])
		}
	}

	return sells[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
