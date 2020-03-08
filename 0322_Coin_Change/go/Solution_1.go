package main

import "math"

// 动态规划，背包问题，背包在内层
// 时间复杂度：O(sn)　　空间复杂度：O(s)　其中，s为总数额,n为硬币种类

func coinChange_1(coins []int, amount int) int {

	if amount==0 {
		return 0
	}

	// dp[i]表示凑出i数值所需要的最少硬币数
	dp := make([]int, amount+1)

	for _, coin := range coins {
		for i:=coin; i<=amount; i++ {
			if i==coin {
				dp[i] = 1
			}else if dp[i-coin]!=0 && dp[i]==0 {
				dp[i] = dp[i-coin] + 1
			}else if dp[i-coin]!=0 && dp[i]!=0 {
				dp[i] = int(math.Min(float64(dp[i-coin] + 1), float64(dp[i])))
			}
		}
	}

	if dp[amount]==0 {
		return -1
	}else {
		return dp[amount]
	}
}