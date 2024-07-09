package main

// 思路:  所有的不被包围的'O'都直接或间接与边界上的'O'相连, 需要从四个边的'O'位置作为起点出发作dfs搜索
// 从内层作为起点出发会有误判

// 时间复杂度: O(m*n)    空间复杂度: O(m*n)

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])

	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	inArea := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	// dfs 寻找'O'位置，并将其标记为'A'
	var dfs func(board [][]byte, x, y int)
	dfs = func(board [][]byte, x, y int) {
		board[x][y] = 'A'

		for k := 0; k < 4; k++ {
			nextX := x + dir[k][0]
			nextY := y + dir[k][1]

			if inArea(nextX, nextY) && board[nextX][nextY] == 'O' {
				dfs(board, nextX, nextY)
			}
		}
	}

	// 分别从四周的'O'位置为起点出发，作dfs搜索
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(board, i, n-1)
		}
	}

	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			dfs(board, 0, i)
		}
		if board[m-1][i] == 'O' {
			dfs(board, m-1, i)
		}
	}

	// 将被标记为'A'的区域改回'O', 未被标记为'A'的'O'区域改未'X'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}
