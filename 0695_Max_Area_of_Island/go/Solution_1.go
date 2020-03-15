package main

// DFS
// 时间复杂度：O(r*c),每个网格最多访问一次  空间复杂度：O(r*c),递归栈深度最多为r*c  (r为行数，c为列数)

var M int  // 行数
var N int  // 列数
var dir = [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}}  // 四个方向

func maxAreaOfIsland_1(grid [][]int) int {

	max_area := 0
	M = len(grid)
	N = len(grid[0])

	for i:=0; i<M; i++ {
		for j:=0; j<N; j++ {
			area := dfs(grid, i, j)
			if area>max_area {
				max_area = area
			}
		}
	}

	return max_area
}

func dfs(grid [][]int, i int, j int) int{

	if grid[i][j]==0 || grid[i][j]==2 {
		return 0
	}

	grid[i][j] = 2  // 避免重复访问岛屿，沉岛思想
	area := 1
	for k:=0; k<4; k++ {
		next_i := i + dir[k][0]
		next_j := j + dir[k][1]

		if inArea(next_i, next_j) {
			area = area + dfs(grid, next_i, next_j)
		}
	}

	return area
}

func inArea(i int, j int) bool {
	return  i>=0 && i<M && j>=0 && j<N 
}