package main

// 思路: 复用第0行、第0列做标记

// 时间复杂度: O(m*n)    空间复杂度: O(1)

func setZeroes_1(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	// 判断第0行、第0列中是否存在零
	row0IsZero, col0IsZero := false, false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0IsZero = true
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			row0IsZero = true
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				// 更新第0行、第0列为零， 同时标记该行、该列后续应全部设置为零
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 根据第0行、第0列是否为零来判断是否更新该行、该列为零
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 最后判断是否将第0行、第0列全部更新为零
	if row0IsZero {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if col0IsZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
