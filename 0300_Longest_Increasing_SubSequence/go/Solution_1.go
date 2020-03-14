package main

import "math"

// 动态规划
// dp[k]表示以第k个元素结尾的数组中的最长上升子序列
// dp[k] = max(dp[k], dp[i]+1) if nums[k]>nums[i], (0<=i<k)
// 时间复杂度：O(n^2)  空间复杂度：O(n)

func lengthOfLIS_1(nums []int) int {

	if len(nums)==0 {
		return 0
	}

	dp := make([]int, len(nums))  // dp[k]表示以第k个元素结尾的数组中的最长上升子序列
	dp[0] = 1
	max_length := dp[0]  // 哨兵变量，记录最长长度

	for k:=1; k<len(nums); k++ {
		dp[k] = 1
		for i:=0; i<k; i++ {
			if nums[k] > nums[i] {
				dp[k] = int(math.Max(float64(dp[k]), float64(dp[i]+1)))
			}
		}

		max_length = int(math.Max(float64(max_length), float64(dp[k])))
	}

	return max_length
}