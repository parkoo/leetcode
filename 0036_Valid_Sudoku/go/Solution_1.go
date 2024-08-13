package main

// 思路：遍历时记录状态
// 小3x3宫格用一维数组表示，坐标映射方式为[i,j] -> i/3*3 + j/3 = i + j/3

// 时间复杂度：O(1)  空间复杂度：O(1) 数独共有 81 个单元格，只需要对每个单元格遍历一次即可

func isValidSudoku_1(board [][]byte) bool {

	checkRow := make([]map[int]int, 9)
	checkCol := make([]map[int]int, 9)
	checkBox := make([]map[int]int, 9)

	for i := 0; i < 9; i++ {
		checkRow[i] = make(map[int]int, 0)
		checkCol[i] = make(map[int]int, 0)
		checkBox[i] = make(map[int]int, 0)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				if _, ok := checkRow[i][num]; ok {
					return false
				} else {
					checkRow[i][num] = 1
				}

				if _, ok := checkCol[j][num]; ok {
					return false
				} else {
					checkCol[j][num] = 1
				}

				// 从左至右 从上至下 3x3宫格的索引为 i*3/3+j/3
				if _, ok := checkBox[i/3*3+j/3][num]; ok {
					return false
				} else {
					checkBox[i/3*3+j/3][num] = 1
				}
			}
		}
	}

	return true
}
