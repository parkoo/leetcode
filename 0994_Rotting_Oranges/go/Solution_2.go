package main

// 思路: 记录层数的多源BFS法，为了记录层数，需要记录并维护每层入队节点的个数'cnt'

// 时间复杂度：O(mn)  空间复杂度：O(mn)

var M_2, N_2 int
var dir_2 = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 方向

func orangesRotting_2(grid [][]int) int {
	M_2 = len(grid)
	N_2 = len(grid[0])

	var q [][2]int
	cnt := 0
	times := 0

	// 多源初始节点(初始所有腐烂的橘子)入队，并记录本层入队节点个数'cnt'
	for i := 0; i < M_2; i++ {
		for j := 0; j < N_2; j++ {
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
				cnt++
			}
		}
	}

	for len(q) > 0 {
		temp := 0 // 中间变量，暂存下一层入队节点个数，本层节点全部出队后，更新给cnt

		// 对本层的cnt个节点出队，并扩展下一层节点，将满足条件的下层节点(新鲜橘子)入队
		// 注意，入队时要更新节点状态(新鲜橘子 -> 腐烂橘子)，避免重复入队
		for i := 0; i < cnt; i++ {
			x := q[0][0]
			y := q[0][1]
			q = q[1:]

			for k := 0; k < 4; k++ {
				next_i := x + dir_2[k][0]
				next_j := y + dir_2[k][1]

				if inArea_2(next_i, next_j) && grid[next_i][next_j] == 1 {
					grid[next_i][next_j] = 2 // 更新节点状态
					q = append(q, [2]int{next_i, next_j})
					temp++
				}
			}
		}
		cnt = temp   // 更新队列中节点个数
		if cnt > 0 { // 若新入队节点大于0,则表明有新橘子被腐烂，层数(也即时间）加一
			times++
		}
	}

	for i := 0; i < M_2; i++ {
		for j := 0; j < N_2; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return times
}

func inArea_2(i int, j int) bool {
	return i >= 0 && i < M_2 && j >= 0 && j < N_2
}
