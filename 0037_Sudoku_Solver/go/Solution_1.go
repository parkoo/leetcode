package main

// 回溯法

func solveSudoku(board [][]byte) {
	rowUsed := make([][]bool, len(board))   // rowUsed[i][v] 表示第i行存在数字v
	colUsed := make([][]bool, len(board))   // rowUsed[i][v] 表示第i列存在数字v
	blockUsed := make([][]bool, len(board)) // rowUsed[i][v] 表示第i个小方框内存在数字v
	for i := 0; i < len(board); i++ {
		rowUsed[i] = make([]bool, len(board)+1)
		colUsed[i] = make([]bool, len(board)+1)
		blockUsed[i] = make([]bool, len(board)+1)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' {
				val := board[i][j] - '0'
				rowUsed[i][val] = true
				colUsed[j][val] = true
				blockUsed[i/3*3+j/3][val] = true
			}
		}
	}

	var backtrack func(board [][]byte, x, y int) bool
	backtrack = func(board [][]byte, x, y int) bool {
		// 自上向下 自左向右 寻找未填充的位置 直至填充完
		for board[x][y] != '.' {
			y++
			if y >= len(board) {
				x++
				y = 0
			}
			if x >= len(board) {
				return true
			}
		}

		for i := 1; i <= 9; i++ {
			if !rowUsed[x][i] && !colUsed[y][i] && !blockUsed[x/3*3+y/3][i] {
				rowUsed[x][i] = true
				colUsed[y][i] = true
				blockUsed[x/3*3+y/3][i] = true
				board[x][y] = '0' + byte(i)

				// 若下一个空位可以填充 则当前空位不需回退
				if backtrack(board, x, y) {
					return true
				} else {
					// 回退
					board[x][y] = '.'
					rowUsed[x][i] = false
					colUsed[y][i] = false
					blockUsed[x/3*3+y/3][i] = false
				}
			}
		}
		return false
	}

	backtrack(board, 0, 0)
}
