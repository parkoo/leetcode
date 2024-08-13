package main

// 思路：遍历时记录状态
// 小3x3宫格用二维数组表示，坐标映射方式为[i,j] -> [i/3, j/3]

// 时间复杂度：O(1)  空间复杂度：O(1) 数独共有 81 个单元格，只需要对每个单元格遍历一次即可

func isValidSudoku_2(board [][]byte) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])

	// 记录第i行、第j列是否出现某个数字
	// rowsUsed[3][8] = true 表示第3行出现过数字8
	rowsUsed := make([]map[byte]bool, m)
	colsUsed := make([]map[byte]bool, m)
	for i := 0; i < m; i++ {
		rowsUsed[i] = make(map[byte]bool)
		colsUsed[i] = make(map[byte]bool)
	}
	// 记录第[i/3, j/3]个3x3宫格是否出现某个数字
	// subBoxUsed[1][2][5] = true 表示第[1,2]个3x3宫格出现过数字5
	subBoxUsed := make([][]map[byte]bool, m/3)
	for i := 0; i < m/3; i++ {
		subBoxUsed[i] = make([]map[byte]bool, m/3)
		for j := 0; j < m/3; j++ {
			subBoxUsed[i][j] = make(map[byte]bool)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			curVal := board[i][j]
			if curVal == '.' {
				continue
			}

			if rowsUsed[i][curVal] || colsUsed[j][curVal] || subBoxUsed[i/3][j/3][curVal] {
				return false
			}

			rowsUsed[i][curVal] = true
			colsUsed[j][curVal] = true
			subBoxUsed[i/3][j/3][curVal] = true
		}
	}

	return true
}
