package main

// 思路: 从右上角或左下角出发，通过大小确定搜索的方向

// 时间复杂度: O(m+n) （在搜索过程中，如果没有找到target, 则要么将y减少1，要么将x增加 1, 最多操作m+n次）   空间复杂度: O(1)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])

	inArea := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	x, y := 0, n-1
	for inArea(x, y) {
		if matrix[x][y] == target {
			return true
		}

		if matrix[x][y] > target {
			y--

		} else {
			x++
		}
	}

	return false
}
