package main

// 思路: 通过翻转实现旋转
// 沿着矩阵左对角线翻转,再沿着矩阵垂直中心线翻转

// 时间复杂度: O(n^2)    空间复杂度: O(1)

func rotate_1(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])

	// 矩阵转置(把行和列互换)，即沿着矩阵左对角线翻转
	for i := 0; i < m; i++ {
		for j := i; j < n; j++ { // 只需要遍历上三角形或下三角形即可
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 对矩阵的每一行做翻转，即沿着矩阵垂直中心线翻转
	for i := 0; i < m; i++ {
		l, r := 0, n-1
		for l < r {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}
}
