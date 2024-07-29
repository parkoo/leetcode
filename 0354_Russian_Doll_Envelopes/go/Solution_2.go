package main

import (
	"sort"
)

// 思路： 贪心 + 二分

// 经过排序处理后 转化为lc0300

// 时间复杂度：O(nlgn)  空间复杂度：O(n)

func maxEnvelopes_2(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}

	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] ||
			(envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})

	d := make([]int, n+1)
	maxLen := 1
	d[maxLen] = envelopes[0][1]

	for i := 1; i < n; i++ {
		curHeigh := envelopes[i][1]
		if curHeigh > d[maxLen] {
			maxLen++
			d[maxLen] = curHeigh

		} else {
			index := binarySearch(d, maxLen, curHeigh)
			d[index] = curHeigh
		}
	}

	return maxLen
}

func binarySearch(arr []int, maxLen int, target int) int {
	left, right := 1, maxLen
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] >= target {
			right = mid

		} else {
			left = mid + 1
		}
	}

	return left
}
