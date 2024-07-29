package main

// 思路: 无后效性 动态规划  类似打家劫舍问题

// 时间复杂度: O(n)    空间复杂度: O(n)

func mostPoints(questions [][]int) int64 {
	n := len(questions)

	// 根据无后效性，可以从后向前遍历
	// dp[i] = k, 表示解决第 i 道题目及以后的题目可以获得的最高分数为k
	// 不解决第 i 道题目，此时 dp[i]=dp[i+1]
	// 解决第 i 道题目, 则，dp[i]=points[i]+dp[i+brainpower[i]+1]
	// 从两者中作选择
	dp := make([]int, n)
	dp[n-1] = questions[n-1][0]

	for i := n - 2; i >= 0; i-- {
		next := 0
		if i+questions[i][1]+1 < n {
			next = dp[i+questions[i][1]+1]
		}

		dp[i] = max(questions[i][0]+next, dp[i+1])
	}

	return int64(dp[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
