package main

import "math"

// 递归 会超时

func superEggDrop(k int, n int) int {
	return solve(k, n)
}

func solve(k, n int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}

	res := math.MaxInt
	for i := 1; i <= n; i++ {
		res = min(res, max(solve(k, n-i), solve(k-1, i-1))+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
