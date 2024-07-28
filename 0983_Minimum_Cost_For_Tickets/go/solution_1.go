package main

// 思路: 动态规划 + 贪心思想

// 时间复杂度: O(endDay)    空间复杂度: O(endDay)

func mincostTickets(days []int, costs []int) int {
	startDay, endDay := days[0], days[len(days)-1]

	// 子问题：dp[i] = k, 表示到第i天结束时，需要的最少花费
	// 转移方程：
	// 如果第 i 天需要通行证，dp[i] = min(dp[i - 1] + cost[0], dp[i - 7] + cost[1], dp[i - 30] + cost[2])
	// 如果第 i 天不需要通行证，dp[i] = dp[i - 1]
	dp := make([]int, endDay+1)

	index := 0
	for i := startDay; i <= endDay; i++ {
		// 第i天是出发日，需要要出发，需要作决策
		if i == days[index] {
			// 购买1天的通行证，需要的花费
			// 在第(i-1)天购买的1天的通行证在第i天已经过期的情况
			x := costs[0]
			if i > 1 {
				x = dp[i-1] + costs[0] // 第(i-1)天结束时的最少花费 + 当前购买1天通行证的花费
			}

			// 购买7天的通行证，需要的花费
			// 在第(i-7)天购买的7天的通行证在第i天已经过期
			y := costs[1]
			if i > 7 {
				y = dp[i-7] + costs[1] // 第(i-7)天结束时的最少花费 + 当前购买7天通行证的花费
			}

			// 购买30天的通行证，需要的花费
			// 在第(i-30)天购买的30天的通行证在第i天已经过期
			z := costs[2]
			if i > 30 {
				z = dp[i-30] + costs[2] // 第(i-30)天结束时的最少花费 + 当前购买30天通行证的花费
			}

			// 从三种选择中做决策
			dp[i] = min(x, min(y, z))

			index++

			// 第i天不需要出发,dp[i]维持在前一天的状态
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[endDay]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
