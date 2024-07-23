package main

// 思路:  降维动态规划

// 时间复杂度: O(mn)    空间复杂度: O(n)

func isInterleave_2(s1 string, s2 string, s3 string) bool {
	m, n, k := len(s1), len(s2), len(s3)
	if m+n != k {
		return false
	}

	// dp[i][j] 表示s1的前i个字符[0 ~ i-1]和s2的前j个字符[0 ~ j-1] 是否可以拼接成s3的前i+j个字符
	// 如：dp[3][4] = true 表示s1[0:2] 和 s2[0:3] 可以拼接出 s3[0:6]
	// 滚动数组降维处理
	dp := make([]bool, n+1)

	// 初始化状态边界
	dp[0] = true
	for j := 1; j <= n; j++ {
		dp[j] = dp[j-1] && s2[j-1] == s3[j-1]
	}

	// dp
	for i := 1; i <= m; i++ {
		// 初始化左边界
		dp[0] = dp[0] && s1[i-1] == s3[i-1]

		for j := 1; j <= n; j++ {
			p := i + j - 1
			dp[j] = (dp[j] && s1[i-1] == s3[p]) || (dp[j-1] && s2[j-1] == s3[p])
		}
	}

	return dp[n]
}
