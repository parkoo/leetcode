package main

// 思路: 同solution_2, 多源BFS

// 时间复杂度：O(mn)  空间复杂度：O(mn)

func orangesRotting_3(grid [][]int) int {
	res := 0

	m := len(grid)
	if m == 0 {
		return res
	}
	n := len(grid[0])

	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	inArea := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	// 多源节点bfs
	bfs := func(gird [][]int, srcs [][2]int) int {
		q := make([][]int, 0)
		// 多个源节点同时入队
		for i := 0; i < len(srcs); i++ {
			q = append(q, []int{srcs[i][0], srcs[i][1]})
			grid[srcs[i][0]][srcs[i][1]] = 2
		}

		times := 0 // 记录扩散次数
		for len(q) > 0 {
			// len(q) 记录了当前层的节点数, 但不能直接用在for的条件中，因为for循环中len(q)在变化
			cnt := len(q)
			for i := 0; i < cnt; i++ {
				cur := q[0]
				x, y := cur[0], cur[1]
				q = q[1:]

				for k := 0; k < 4; k++ {
					nextX := x + dir[k][0]
					nextY := y + dir[k][1]

					if inArea(nextX, nextY) && grid[nextX][nextY] == 1 {
						q = append(q, []int{nextX, nextY})
						grid[nextX][nextY] = 2
					}
				}
			}

			// 引入下一层节点，则表示向下扩散一轮
			if len(q) > 0 {
				times++
			}
		}

		return times
	}

	// 寻找所有源节点
	srcs := make([][2]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				srcs = append(srcs, [2]int{i, j})
			}
		}
	}

	// bfs扩散
	res = bfs(grid, srcs)

	// 经过bfs扩散后，判断是否还有没感染的个体
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return res
}
