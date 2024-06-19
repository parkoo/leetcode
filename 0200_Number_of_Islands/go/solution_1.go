package main

// 思路: DFS遍历图

// 时间复杂度: O(m*n)    空间复杂度: O(m*n)

func numIslands_1(grid [][]byte) int {
	res := 0
	if len(grid) == 0 {
		return res
	}

	m, n := len(grid), len(grid[0])

	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	inArea := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	var dfs func(grid [][]byte, x, y int)
	dfs = func(grid [][]byte, x, y int) {
		used[x][y] = true

		for k := 0; k < 4; k++ {
			nextX := x + dir[k][0]
			nextY := y + dir[k][1]

			if inArea(nextX, nextY) && !used[nextX][nextY] && grid[nextX][nextY] == '1' {
				dfs(grid, nextX, nextY)
			}
		}
	}

	// 遍历图
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !used[i][j] && grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}

	return res
}
