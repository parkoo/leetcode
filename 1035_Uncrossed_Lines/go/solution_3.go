package main

// 思路: 动态规划 在solution_2的基础上降维空间 解法类似lc1143

// 时间复杂度: O(mn)    空间复杂度: O(n)

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)

	// dp[i][j] = k, 表示以nums1[i-1]、nums2[j-1]结尾的最多不想交连线的条数为k
	dp := make([]int, n+1)

	for i := 1; i <= m; i++ {
		pre := 0 // dp[i-1][j-1]
		for j := 1; j <= n; j++ {
			temp := dp[j]
			if nums1[i-1] == nums2[j-1] {
				dp[j] = pre + 1

			} else {
				dp[j] = max(dp[j], dp[j-1])
			}

			pre = temp
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
