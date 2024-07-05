package main

import "sort"

// 思路:  排序 + 贪心（尽可能使得一个重叠区域可以引爆的气球数最大） 重点在于怎么去维护这个重叠区域

// 时间复杂度: O(nlgn)    空间复杂度: O(1)

func findMinArrowShots(points [][]int) int {
	res := 0

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || (points[i][0] == points[j][0] && points[i][1] < points[j][1])
	})

	// public代表相交区域
	// 若新来的区间能够与当前public重叠，则将新来的区间划分给当前public, 同时更新当前public为当前public与新来区间重叠的部分
	// 若新来的区间不能够与当前public重叠， 则将当前public作为一组直接射爆，并将新来的区间作为新的public
	public := points[0]
	res++
	for i := 1; i < len(points); i++ {
		if isOverload(public, points[i]) {
			public = getOverload(public, points[i])
			continue
		}

		public = points[i]
		res++
	}

	return res
}

func isOverload(a, b []int) bool {
	if a[1] < b[0] || a[0] > b[1] {
		return false
	}

	return true
}

func getOverload(a, b []int) []int {
	start, end := a[0], a[1]
	if b[0] > a[0] {
		start = b[0]
	}
	if b[1] < a[1] {
		end = b[1]
	}

	return []int{start, end}
}
