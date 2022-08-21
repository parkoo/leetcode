package main

// 动态规划  dp[i][j] = m 表示做 i 次操作，而且有 k 个鸡蛋，最多可以能找到的最高层数为 m
// 时间复杂度：O(kn)  空间复杂度：O(kn)  共有kn个子问题(子状态)

func superEggDrop(k int, n int) int {
	if n == 1 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		if i == 1 {
			for j := 1; j <= k; j++ {
				dp[i][j] = 1
			}
		}
	}

	res := -1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j] + 1
		}

		if dp[i][k] >= n {
			res = i
			break

		}

	}

	return res
}
