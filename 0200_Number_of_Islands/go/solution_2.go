package main

// 思路: BFS遍历图

// 时间复杂度: O(m*n)    空间复杂度: O(m*n)

func numIslands_2(grid [][]byte) int {
	res := 0
	if len(grid) == 0 {
		return res
	}

	m, n := len(grid), len(grid[0])

	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	inArea := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	bfs := func(grid [][]byte, x, y int) {
		q := make([][]int, 0)
		q = append(q, []int{x, y})
		used[x][y] = true // 入队后标记为已访问

		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			x, y := cur[0], cur[1]

			for k := 0; k < 4; k++ {
				nextX := x + dir[k][0]
				nextY := y + dir[k][1]
				if inArea(nextX, nextY) && grid[nextX][nextY] == '1' && !used[nextX][nextY] {
					q = append(q, []int{nextX, nextY})
					used[nextX][nextY] = true // 入队后标记为已访问
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !used[i][j] {
				res++
				bfs(grid, i, j)
			}
		}
	}

	return res
}
