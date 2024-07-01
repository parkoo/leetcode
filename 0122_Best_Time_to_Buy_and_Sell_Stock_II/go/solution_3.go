package main

// 贪心法  因为交易次数不受限，如果可以把所有的上坡全部收集到，一定是利益最大化的
// 时间复杂度：O(n)  空间复杂度：O(1)

func maxProfit_3(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
