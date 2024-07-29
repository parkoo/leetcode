package main

import (
	"sort"
)

// 思路： 动态规划

// 信封的宽为w, 高为h
// 首先我们将所有的信封按照 w 值第一关键字升序、h 值第二关键字降序进行排序
// 按照上述方式排序后，我们就可以忽略 w 维度
// 因为 对于 i > j, 一定有 w[i] >= w[j] w不在是限制只需要考虑h， 而当w[i] == w[j]时，h[i] < h[j], 即排除了将宽度相等的信封进行嵌套的情况
// 因此，最终求出 h 维度的最长严格递增子序列，其长度即为答案

// 经过排序后可以转化为对信封高度h求最长递增子序列问题 即同lc0300

// 时间复杂度：O(n^2)  空间复杂度：O(n)

func maxEnvelopes_1(envelopes [][]int) int {

	// 排序
	sort.Slice(envelopes, func(i int, j int) bool {
		return envelopes[i][0] < envelopes[j][0] ||
			(envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
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
