package main

// 哨兵法
// 时间复杂度：O(n)  空间复杂度：O(1)

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	// 设置两个哨兵
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < n; i++ {
		maxProfit = max(maxProfit, prices[i]-minPrice)
		minPrice = min(prices[i], minPrice)
	}

	return maxProfit
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
