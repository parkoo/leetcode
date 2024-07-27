package main

// 思路 ：完全背包最值问题，物品在外层背包在内层， 更好的解法见solution_5, 先设置好初值方便处理

// 时间复杂度：O(sn)　　空间复杂度：O(s)　其中，s为总数额,n为硬币种类

func coinChange_1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// dp[i]表示凑出i数值所需要的最少硬币数
	dp := make([]int, amount+1)

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if i == coin {
				dp[i] = 1
			} else if dp[i-coin] != 0 && dp[i] == 0 {
				dp[i] = dp[i-coin] + 1
			} else if dp[i-coin] != 0 && dp[i] != 0 {
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}

	if dp[amount] == 0 {
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
