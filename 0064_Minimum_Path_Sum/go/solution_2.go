package main

// 思路: 动态规划 一维空间
// 定义状态 dp[i][j] 表示到达坐标（i,j）时的最小路径和， 到达（i，j）位置时，可能从上方（i-1, j）过来或者从左方(i, j-1)过来
// 状态转移方程：dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
// 由于当前状态只依赖上方和左方的状态（最左边一列只依赖上方状态），所以，只需要设置一维空间即可, 由于当前转台依赖当前行的左边的值，所以第二层循环从左向右遍历（需要先计算出左边的值）

// 时间复杂度: O(m*n)    空间复杂度: O(n)

func minPathSum_2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n)

	// 设置初始状态
	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		// 设置最左边一列的初始状态，仅依赖上方的状态
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
		}
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
