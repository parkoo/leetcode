package main

import "math"

// 简单的动态规划，只需记录两个状态
// 时间复杂度：O(n)  空间复杂度：O(1)

func maxProfit_1(prices []int) int {

	if len(prices)==0 {
		return 0
	}

	max_profit := 0  //　记录当天截止可获得的最大利润
	min_price := prices[0]  // 记录当天之前股票的最低价格

	for i:=0; i<len(prices); i++ {
		max_profit = int(math.Max(float64(max_profit), float64(prices[i]-min_price)))
		min_price = int(math.Min(float64(prices[i]), float64(min_price)))
	}

	return max_profit
}