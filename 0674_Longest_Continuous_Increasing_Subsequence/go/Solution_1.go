package main

// 动态规划
// 时间复杂度：O(n)  空间复杂度：O(n)

func findLengthOfLCIS_1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1

	res := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
