package main

// 思路: 子序列问题
// 动态规划不是最优解
// 因为差值一定，可以用map存储之前的结果，然后直接查找即可

// 时间复杂度: O(n^2)    空间复杂度: O(n^2)

func longestSubsequence_1(arr []int, difference int) int {
	res := 1
	n := len(arr)

	// dp[i]=k 表示以arr[i]结尾的最长等差子序列的长度为k
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i]-arr[j] == difference {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
