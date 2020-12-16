package main

import "math"

// 动态规划
// 时间复杂度：O(m*n*n), 空间复杂度：O(m*m) m为key的长度，n为ring的长度

func findRotateSteps_1(ring string, key string) int {
	m := len(key)
	n := len(ring)

	// pos[x]记录字符“x”在ring中出现的位置，x=key[i]
	pos := make(map[byte][]int)
	for i := 0; i < len(key); i++ {
		for j := 0; j < len(ring); j++ {
			if ring[j] == key[i] {
				pos[key[i]] = append(pos[key[i]], j)
			}
		}
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化状态
	for _, p := range pos[key[0]] {
		dp[0][p] = dist(n, 0, p) + 1
	}

	// 状态转移
	for i := 1; i < m; i++ {
		for _, p := range pos[key[i]] {
			dp[i][p] = math.MaxInt64
			for _, q := range pos[key[i-1]] {
				dp[i][p] = min(dp[i][p], (dp[i-1][q] + dist(n, p, q) + 1))
			}
		}
	}

	res := math.MaxInt64
	for _, p := range pos[key[m-1]] {
		res = min(res, (dp[m-1][p]))
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 计算环中两个位置的最短距离
func dist(n int, k int, l int) int {
	return min((n+k-l)%n, (n+l-k)%n)
}
