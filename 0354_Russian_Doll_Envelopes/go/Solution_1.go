package main

import (
	"sort"
)

// 动态规划 二维最长递增子序列问题
// 时间复杂度：O(n^2)  空间复杂度：O(n)

func maxEnvelopes_1(envelopes [][]int) int {
	// 排序
	sort.Slice(envelopes, func(i int, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})

	n := len(envelopes)
	dp := make([]int, n)
	maxLen := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
