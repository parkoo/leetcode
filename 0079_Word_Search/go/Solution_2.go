package main

// 回溯法

func exist_2(board [][]byte, word string) bool {
	dir := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
	used := make([][]bool, len(board))
	for i, _ := range used {
		used[i] = make([]bool, len(board[0]))
	}

	var inArea func(x, y int) bool
	inArea = func(x, y int) bool {
		return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
	}

	var backtrack func(start, x, y int, board [][]byte, word string) bool
	backtrack = func(start, x, y int, board [][]byte, word string) bool {
		if board[x][y] != []byte(word[start : start+1])[0] {
			return false
		}

		if start == len(word)-1 {
			return true
		}

		used[x][y] = true
		for k := 0; k < 4; k++ {
			nextX := x + dir[k][0]
			nextY := y + dir[k][1]

			if inArea(nextX, nextY) && !used[nextX][nextY] && backtrack(start+1, nextX, nextY, board, word) {
				return true
			}
		}

		used[x][y] = false
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			res := backtrack(0, i, j, board, word)
			if res {
				return res
			}
		}
	}

	return false
}
