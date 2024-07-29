package main

// 思路: 动态规划 解法类似lc1143
// 左边界上边界设置虚拟区间

// 时间复杂度: O(mn)    空间复杂度: O(mn)

func maxUncrossedLines_2(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)

	// dp[i][j] = k, 表示以nums1[i-1]、nums2[j-1]结尾的最多不想交连线的条数为k
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1

			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
