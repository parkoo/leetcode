package main

// 思路: 动态规划 需要找到一定规律
// 在solution_1的基础上对空间维度降低
// https://leetcode.cn/problems/domino-and-tromino-tiling/solutions/1968516/by-endlesscheng-umpp/?envType=problem-list-v2&envId=D63oKgiL&status=TO_DO&difficulty=MEDIUM

// 时间复杂度: O(n)    空间复杂度: O(1)

func numTilings_2(n int) int {
	if n < 3 {
		return n
	}

	// dp[i]=k 表示2*i的区域的总铺法
	pre1, pre2, pre3 := 1, 1, 2

	res := 0
	for i := 3; i <= n; i++ {
		res = (2*pre3 + pre1) % (1e9 + 7)
		pre1 = pre2
		pre2 = pre3
		pre3 = res
	}

	return res
}
