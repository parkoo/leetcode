package main

// 思路：一维动态规划  每一天的状态只与前一天的状态有关，而与更早的状态都无关，因此我们不必存储这些无关的状态.
// 思路同lc122基本相似

// 时间复杂度：O(n)  空间复杂度：O(1)

func maxProfit_2(prices []int) int {
	dp0 := 0
	dp1 := -prices[0]

	for i := 1; i < len(prices); i++ {
		newDp0 := max(dp0, dp1+prices[i])
		newDp1 := max(dp1, dp0-prices[i])
		dp0 = newDp0
		dp1 = newDp1
	}

	return dp0
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
