package main

// 思路:  对每个单词都进行 DFS 时间复杂度较高 会超时

// 时间复杂度: O(?)    空间复杂度: O(?)

func findWords_1(board [][]byte, words []string) []string {
	res := make([]string, 0)

	if len(board) == 0 || len(board[0]) == 0 {
		return res
	}

	m, n := len(board), len(board[0])
	dir := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	inArea := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	var dfs func(board [][]byte, x, y int, i int, target string) bool
	dfs = func(board [][]byte, x, y int, i int, target string) bool {
		if i == len(target)-1 {
			if board[x][y] == target[i] {
				return true
			}
			return false
		}
		if board[x][y] != target[i] {
			return false
		}

		used[x][y] = true
		for k := 0; k < 4; k++ {
			nextX := x + dir[k][0]
			nextY := y + dir[k][1]
			if inArea(nextX, nextY) && !used[nextX][nextY] && dfs(board, nextX, nextY, i+1, target) {
				return true
			}
		}
		used[x][y] = false

		return false
	}

	for _, word := range words {
		start := word[0]
		used = make([][]bool, m)
		for i := 0; i < m; i++ {
			used[i] = make([]bool, n)
		}

	r:
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if board[i][j] == start {
					if dfs(board, i, j, 0, word) {
						res = append(res, word)
						break r
					}

				}
			}
		}
	}

	return res
}
