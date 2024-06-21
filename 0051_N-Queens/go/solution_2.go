package main

// 思路：同solution_1  回溯法  二维组合问题

// 时间复杂度：O(n!) (第1行尝试n次，第2行尝试（n-1）次... ) 空间复杂度：O(n)

func solveNQueens_2(n int) [][]string {
	res := make([][]string, 0)

	// 同行、同列、同左斜线、同右斜线是否被标记
	rowVisited, colVisited := make([]bool, n), make([]bool, n)
	diagRightVisited, diagLeftVisited := make([]bool, 2*n-1), make([]bool, 2*n-1)

	// 构建棋盘并初始化
	grid := make([][]string, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]string, n)
		for j := 0; j < n; j++ {
			grid[i][j] = "."
		}
	}

	// 定义回溯函数
	var backtracking func(grid [][]string, m int)
	backtracking = func(grid [][]string, m int) {
		// 递归终止
		if m == n {
			// 此时，grid是一个合法解，将该解放入结果集中
			item := make([]string, 0)
			for i := 0; i < n; i++ {
				row := ""
				for j := 0; j < n; j++ {
					row += grid[i][j]
				}
				item = append(item, row)
			}
			res = append(res, item)
			return
		}

		// 尝试在grid[m][i]位置放“皇后”
		for i := 0; i < n; i++ {
			if !rowVisited[m] && !colVisited[i] && !diagLeftVisited[m+i] && !diagRightVisited[m-i+n-1] {
				grid[m][i] = "Q"
				rowVisited[m], colVisited[i] = true, true
				diagLeftVisited[m+i], diagRightVisited[m-i+n-1] = true, true

				backtracking(grid, m+1)

				// 回溯恢复状态
				grid[m][i] = "."
				rowVisited[m], colVisited[i] = false, false
				diagLeftVisited[m+i], diagRightVisited[m-i+n-1] = false, false
			}
		}
	}

	// 回溯
	backtracking(grid, 0)

	return res
}
