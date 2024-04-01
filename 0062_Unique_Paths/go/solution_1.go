package main

// 思路: 动态规划 二维空间
// 定义状态 dp[i][j] 表示机器人到达坐标 (i, j）时，可能走过的不同路径数, 到达该坐标时，可能从上方（i-1, j）过来或者从左方(i, j-1)过来
// 所以，dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 注意设置状态的初值

// 时间复杂度: O(m*n)    空间复杂度: O(m*n)

func uniquePaths_1(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 设置初始状态
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
