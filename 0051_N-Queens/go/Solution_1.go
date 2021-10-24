package main

// 回溯法  二维组合问题
// 时间复杂度：O(n!) (第1行尝试n次，第2行尝试（n-1）次... ) 空间复杂度：O(n)

var result [][]string
var rowUsed, colUsed, diag1Used, diag2Used []bool

func solveNQueens_1(n int) [][]string {
	result = [][]string{}
	rowUsed = make([]bool, n)
	colUsed = make([]bool, n)
	diag1Used = make([]bool, 2*n-1)
	diag2Used = make([]bool, 2*n-1)

	grid := make([][]string, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]string, n)
		for j := 0; j < n; j++ {
			grid[i][j] = "."
		}
	}

	backtrack(grid, 0)
	return result
}

func backtrack(grid [][]string, k int) {

	if k == len(grid) {
		var list []string
		for i := 0; i < len(grid); i++ {
			s := ""
			for j := 0; j < len(grid); j++ {
				s += grid[i][j]
			}
			list = append(list, s)
		}
		result = append(result, list)
		return
	}

	// 	尝试将 第k行 第i列 设置为Queen
	for i := 0; i < len(grid); i++ {
		if !rowUsed[k] && !colUsed[i] && !diag1Used[i+k] && !diag2Used[i-k+len(grid)-1] {
			grid[k][i] = "Q"
			rowUsed[k] = true
			colUsed[i] = true
			diag1Used[i+k] = true
			diag2Used[i-k+len(grid)-1] = true
			backtrack(grid, k+1)
			grid[k][i] = "."
			rowUsed[k] = false
			colUsed[i] = false
			diag1Used[i+k] = false
			diag2Used[i-k+len(grid)-1] = false
		}
	}
}
