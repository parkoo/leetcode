package main

import (
	"sort"
)

// 时间复杂度：O(nlgn)  空间复杂度：O(lgn)(排序中的递归)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1]
	})

	n := len(intervals)
	res := make([][]int, 0)
	res = append(res, intervals[0])
	for i := 1; i < n; i++ {
		if intervals[i][0] > res[len(res)-1][1] || intervals[i][1] < res[len(res)-1][0] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][0] = min(res[len(res)-1][0], intervals[i][0])
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
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

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}
