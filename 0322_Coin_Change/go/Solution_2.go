package main

// 动态规划，完全背包最值问题, 背包在外层物品在内层, 更好的解法见solution_3 ，先设置好初值方便处理

// 同lc0279

// 时间复杂度：O(sn)　　空间复杂度：O(s)　其中，s为总数额,n为硬币种类

func coinChange_2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// dp[i]表示凑出i数值所需要的最少硬币数
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i == coin {
				dp[i] = 1
			} else if i > coin && dp[i-coin] != 0 && dp[i] == 0 {
				dp[i] = dp[i-coin] + 1
			} else if i > coin && dp[i-coin] != 0 && dp[i] != 0 {
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
