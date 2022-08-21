package main

import (
	"fmt"
	"math"
)

// 递归 + 记忆化搜索 + 二分查找
// 时间复杂度：O(knlgn)  空间复杂度：O(kn)  共有kn个子问题(子状态)

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

	memoValue := getMemo(k, n)
	if memoValue != -1 {
		return memoValue
	}

	res := math.MaxInt
	l, r := 1, n
	for l < r {
		m := l + (r-l)/2
		s1 := solve(k, n-m)
		s2 := solve(k-1, m-1)
		if s2 >= s1 {
			r = m
		} else {
			l = m + 1
		}
	}
	res = min(res, max(solve(k, n-l), solve(k-1, l-1))+1)
	if l-1 > 0 {
		res = min(res, max(solve(k, n-(l-1)), solve(k-1, (l-1)-1))+1)
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
