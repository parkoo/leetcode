package main

// 回溯法

func exist_1(board [][]byte, word string) bool {
	res := false
	dir := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	used := make([][]bool, len(board))
	for i, _ := range used {
		used[i] = make([]bool, len(board[0]))
	}

	var inArea func(x, y int) bool
	inArea = func(x, y int) bool {
		return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
	}

	var backtrack func(board [][]byte, x, y, start int, word string)
	backtrack = func(board [][]byte, x, y, start int, word string) {
		if board[x][y] != []byte(word[start : start+1])[0] {
			return
		}

		if start == len(word)-1 {
			res = true
			return
		}

		used[x][y] = true
		for k := 0; k < 4; k++ {
			nextX := x + dir[k][1]
			nextY := y + dir[k][0]

			if inArea(nextX, nextY) && !used[nextX][nextY] {
				backtrack(board, nextX, nextY, start+1, word)
			}
		}
		used[x][y] = false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			backtrack(board, i, j, 0, word)
			if res {
				return res
			}
		}
	}

	return res
}
