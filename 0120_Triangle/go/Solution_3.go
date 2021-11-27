package main

// 原地动态规划
// 时间复杂度：O(n^2)  空间复杂度：O(1)

func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	for i := 1; i < n; i++ {
		triangle[i][0] = triangle[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			triangle[i][j] = min(triangle[i-1][j], triangle[i-1][j-1]) + triangle[i][j]
		}
		triangle[i][i] = triangle[i-1][i-1] + triangle[i][i]
	}

	res := triangle[n-1][0]
	for i := 0; i < n; i++ {
		if triangle[n-1][i] < res {
			res = triangle[n-1][i]
		}
	}

	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
