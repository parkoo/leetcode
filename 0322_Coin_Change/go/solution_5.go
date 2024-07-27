package main

// 动态规划，完全背包最值问题，物品在外层背包在内层
// 注意设置好初值

// 注意和lc0474区分, 其为0-1背包最值问题，需物品在外层且内层从大到小遍历，以保证每个物品只使用一次
// 本题是完全背包最值问题，物品可以在外层也可以在内层，物品在外层时，内层从小到大遍历，保证物品可以重复使用

// 时间复杂度：O(sn)　　空间复杂度：O(s)　其中，s为总数额,n为硬币种类

func coinChange_5(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
