package main

// 思路: 子序列问题  维护以nums[i]结尾的最长子序列的个数

// 时间复杂度: O(n^2)    空间复杂度: O(n)

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	res := 0

	maxLen := 0 // 记录当前全局递增子序列的最长长度

	// 以同一个数字结尾的最长递增子序列的个数可能有多个
	// 如[1,5,4,7], 以’7‘结尾的最长递增子序列有两个[1,5,7]、[1,4,7]
	// maxLenCnt 维护以nums[i]结尾的最长子序列的个数
	// maxLenCnt[i] = k, 表示以nums[i]结尾的最长子序列的个数为k
	maxLenCnt := make(map[int]int)
	// dp[i] = k, 表示以nums[i]结尾的最长子序列的长度为k
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		// 初始化
		dp[i] = 1
		maxLenCnt[i] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {

				// 维护dp[i]、maxLenCnt[i]
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					maxLenCnt[i] = maxLenCnt[j]

				} else if dp[j]+1 == dp[i] {
					maxLenCnt[i] += maxLenCnt[j]
				}
			}
		}

		// 维护maxLen、 res
		if dp[i] > maxLen {
			maxLen = dp[i]
			res = maxLenCnt[i]

		} else if dp[i] == maxLen {
			res += maxLenCnt[i]
		}
	}

	return res
}
