package main

// 思路: 动态规划 需要找到一定规律
// https://leetcode.cn/problems/domino-and-tromino-tiling/solutions/1968516/by-endlesscheng-umpp/?envType=problem-list-v2&envId=D63oKgiL&status=TO_DO&difficulty=MEDIUM

// 时间复杂度: O(n)    空间复杂度: O(n)

func numTilings_1(n int) int {
	if n < 3 {
		return n
	}

	// dp[i]=k 表示2*i的区域的总铺法
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % (1e9 + 7)
	}

	return dp[n]
}
