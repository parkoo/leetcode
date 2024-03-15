package main

// 动态规划之一  状态定义: dp[i]表示"将第i家打劫后", 此时在第i家这个位置可以获得的最大金币, dp[i]不表示全局最大值
// 时间复杂度：O(n^2)  空间复杂度：O(n)

func rob_1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]

	res := dp[0]
	for i := 1; i < n; i++ {
		maxPre := 0
		for j := 0; j < i-1; j++ {
			if dp[j] > maxPre {
				maxPre = dp[j]
			}
		}
		dp[i] = maxPre + nums[i]
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
