package main

// 思路：回溯法 回溯函数带返回值

func exist_2(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	inArea := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	var backtract func(board [][]byte, x, y int, i int) bool
	backtract = func(board [][]byte, x, y int, i int) bool {
		if i == len(word)-1 && board[x][y] == word[i] {
			return true
		}

		if board[x][y] != word[i] {
			return false
		}

		used[x][y] = true
		for k := 0; k < 4; k++ {
			nextX := x + dir[k][0]
			nextY := y + dir[k][1]

			if inArea(nextX, nextY) && !used[nextX][nextY] {
				if backtract(board, nextX, nextY, i+1) {
					return true
				}
			}
		}
		used[x][y] = false

		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtract(board, i, j, 0) {
				return true
			}
		}
	}

	return false
}
