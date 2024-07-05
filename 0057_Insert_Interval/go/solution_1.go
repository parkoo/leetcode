package main

// 思路: 模拟法

// 时间复杂度: O(n)    空间复杂度: O(1)

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)

	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}

	merged := false // 记录newInterval是否已经入队
	for i := 0; i < len(intervals); i++ {
		cur := intervals[i]
		if isOverload(cur, newInterval) { // 与newInterVal有重叠，则合并到newInterval中，最后跟着newInterval入队
			newInterval = merge(cur, newInterval)
			continue
		}

		if cur[0] > newInterval[1] && !merged {
			res = append(res, newInterval)
			merged = true
		}

		res = append(res, cur)
	}

	if !merged {
		res = append(res, newInterval)
	}

	return res
}

func isOverload(a, b []int) bool {
	if b[0] > a[1] || b[1] < a[0] {
		return false
	}

	return true
}

func merge(a, b []int) []int {
	start, end := a[0], a[1]
	if b[0] < a[0] {
		start = b[0]
	}
	if b[1] > a[1] {
		end = b[1]
	}

	return []int{start, end}
}
