package main

// 思路：完全背包最值问题, 背包在外层物品在内层
// 注意设置初值
// 定义 dp[i] 表示凑出i数值所需要的最少硬币数
// 状态转移：dp[i] = min(dp[i], dp[i-coins[j]] + 1) when 0<=j<=len(coins)

// 同lc0279

// 时间复杂度：O(sn)　　空间复杂度：O(s)　其中，s为总数额,n为硬币种类

func coinChange_3(coins []int, amount int) int {
	dp := make([]int, amount+1)

	initVal := amount + 1 // 定义一个不可能的且大于可能结果的初始值
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = initVal // 遍历的过程中同时作初始化
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	res := dp[amount]
	if res == initVal {
		res = -1
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
