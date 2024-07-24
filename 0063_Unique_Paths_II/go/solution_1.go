package main

// 思路: 降维动态规划

// 时间复杂度: O(mn)    空间复杂度: O(n)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 { // 起点是障碍物
		return 0
	}

	// dp[i][j] 表示从起点(0,0)到(i,j)的路径数
	// 如果当前位置不是障碍物, dp[i][j] = dp[i-1][j] + dp[i][j-1]
	// 如果当前位置是障碍物, dp[i][j] = 0
	// 滚动数组降维处理
	dp := make([]int, n)

	dp[0] = 1
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[j] = dp[j-1]
		}
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[0] = dp[0]
		} else {
			dp[0] = 0
		}

		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[j] = dp[j] + dp[j-1]
			} else {
				dp[j] = 0
			}
		}
	}

	return dp[n-1]
}
