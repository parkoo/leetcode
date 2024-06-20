package main

// 思路: 模拟法

// 时间复杂度: O(m*n)    空间复杂度: O(1)

func spiralOrder_1(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}

	// left: 维护左边界 right:维护右边界 top:维护上边界 bottom:维护下边界
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1

	for left <= right && top <= bottom {
		// 从左向右遍历, 注意判断上下边界
		for i := left; i <= right && top <= bottom; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		// 从上向下遍历, 注意判断左右边界
		for i := top; i <= bottom && left <= right; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		// 从右向左遍历, 注意判断上下边界
		for i := right; i >= left && top <= bottom; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--

		// 从下向上遍历，注意判断左右边界
		for i := bottom; i >= top && left <= right; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res
}
