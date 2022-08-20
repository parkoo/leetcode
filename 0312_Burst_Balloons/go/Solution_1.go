package main

// 动态规划 dp[i][i] = x 表示戳破气球i和气球j之间（开区间，不包括i和j）的所有气球，可以获得的最高分数为x
// 时间复杂度：O(n^3)  空间复杂度：O(n^2)

func maxCoins(nums []int) int {
	n := len(nums)
	nums = append([]int{1}, append(nums, 1)...)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for i := n; i >= 0; i-- {
		for j := i + 1; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][k]+dp[k][j]+nums[i]*nums[k]*nums[j], dp[i][j])
			}
		}
	}

	return dp[0][n+1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
