package main

// 思路:  动态规划  子序列问题

// 时间复杂度: O(n^2)    空间复杂度: O(n*diff)

func longestArithSeqLength_1(nums []int) int {
	res := 0
	n := len(nums)
	if n == 0 {
		return res
	}

	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}
	maxDiff := maxVal - minVal

	// dp[i][k] = x, 表示以nums[i]结尾的公差为k的等差序列的最长长度为x
	dp := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
	}

	// 设置dp[0]的初值
	for i := -maxDiff; i <= maxDiff; i++ {
		dp[0][i] = 1
	}
	res = 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			// 懒设置dp[j][diff]的初值
			if _, ok := dp[j][diff]; !ok {
				dp[j][diff] = 1
			}

			dp[i][diff] = max(dp[i][diff], dp[j][diff]+1)
			res = max(res, dp[i][diff])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
