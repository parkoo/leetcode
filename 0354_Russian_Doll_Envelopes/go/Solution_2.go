package main

import (
	"sort"
)

// 贪心 + 二分
// 时间复杂度：O(nlgn)  空间复杂度：O(n)

func maxEnvelopes_2(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})

	n := len(envelopes)
	// 用len(dp)表示最长的递增子序列
	// dp 索引从0开始，取dp[i]时，len(dp)==i+1, 以dp[i] 记录长度为i+1的递增子序列的末尾元素的最小值
	dp := make([]int, 0)
	dp = append(dp, envelopes[0][1])

	for i := 1; i < n; i++ {
		e := envelopes[i][1]
		if e > dp[len(dp)-1] {
			dp = append(dp, e)
		} else {
			// 若dp中不存在元素e, 则返回元素e应该插入的位置
			k := sort.SearchInts(dp, e)
			dp[k] = e
		}
	}

	return len(dp)
}
