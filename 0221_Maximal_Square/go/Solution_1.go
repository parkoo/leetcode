package main

// 思路：动态规划

// 时间复杂度；O(mn)  空间复杂度：O(mn)

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	// dp[i][j] 表示以坐标点(i, j)为右下角的值为'1'最大正方形的边长
	dp := make([][]int, m+1) // 虚拟边界
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1) // 虚拟边界
	}

	maxLen := 0 // 记录最长边长
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 虚拟边界，dp数组向左向上多一层边界
			if matrix[i-1][j-1] == '1' {
				// 取短板
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}

			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}

	return maxLen * maxLen
}

func min(i, j, k int) int {
	min := i
	if j < min {
		min = j
	}
	if k < min {
		min = k
	}

	return min
}
