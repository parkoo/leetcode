package main

// 动态规划
// 时间复杂度：O(n^2)  空间复杂度：O(n)

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	// dp[i]表示跳到第i个位置所需的最少步数
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = 1<<31 - 1
	}

	for i := 0; i < n; i++ {
		// dp[i]表示跳到第i个位置所需的最少步数，如果此时可以直接跳到终点，则直接可得跳到终点的最少步数
		if nums[i] >= n-1-i {
			return dp[i] + 1
		}

		// 从当前位置向后推，向后更新状态
		for step := 0; step <= nums[i]; step++ {
			if i+step < n && dp[i]+1 < dp[i+step] {
				dp[i+step] = dp[i] + 1
			}
		}
	}

	return dp[n-1]
}
