package main

import "sort"

// 思路: 动态规划
// 类似lc0300

// 时间复杂度: O(n^2)    空间复杂度: O(n)

func findLongestChain(pairs [][]int) int {
	n := len(pairs)
	if n < 2 {
		return n
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0] || (pairs[i][0] == pairs[j][0] && pairs[i][1] < pairs[j][1])
	})

	// dp[i] = k, 表示以pairs[i]结尾的最长数对链的长度为k
	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
