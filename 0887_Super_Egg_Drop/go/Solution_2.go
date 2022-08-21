package main

import (
	"fmt"
	"math"
)

// 递归 + 记忆化搜索 会超时
// 时间复杂度：O(kn^2)  空间复杂度：O(kn)  共有kn个子问题(子状态)

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
	setMemo(k, n, res)
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

var memo = map[string]int{}

func getMemo(k, n int) int {
	key := fmt.Sprintf("%d_%d", k, n)
	if v, ok := memo[key]; ok {
		return v
	}
	return -1
}

func setMemo(k, n, v int) {
	key := fmt.Sprintf("%d_%d", k, n)
	memo[key] = v
}
