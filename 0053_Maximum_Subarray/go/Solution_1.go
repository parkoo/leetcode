package main

// 动态规划
// 时间复杂度：O(n)  空间复杂度：O(n)

func maxSubArray_1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// dp[i]表示以第i个位置为子序列结尾时，可以得到的最大子序和
	dp := make([]int, n)
	dp[0] = nums[0]

	maxSum := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxSum = max(maxSum, dp[i])
	}

	return maxSum
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
