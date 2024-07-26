package main

// 思路: 动态规划

// 时间复杂度: O(n)    空间复杂度: O(1)

func minCostClimbingStairs(cost []int) int {
	pre1, pre2 := 0, 0
	res := 0
	for i := 2; i <= len(cost); i++ {
		res = min(cost[i-2]+pre1, cost[i-1]+pre2)

		pre1 = pre2
		pre2 = res
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
