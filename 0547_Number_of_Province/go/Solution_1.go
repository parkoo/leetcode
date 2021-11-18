package main

// dfs 标记法
// 时间复杂度：O(n^2)  空间复杂度：O(n)

func findCircleNum(isConnected [][]int) int {
	res := 0
	n := len(isConnected)
	used := make([]bool, n)

	var dfs func(isConnected [][]int, i int)
	dfs = func(isConnected [][]int, i int) {
		for k := 0; k < n; k++ {
			// 有连接关系且没有被标记
			if isConnected[i][k] == 1 && !used[k] {
				used[k] = true
				dfs(isConnected, k)
			}
		}
	}

	for i := 0; i < n; i++ {
		// 对没有标记的点进行DFS
		if !used[i] {
			dfs(isConnected, i)
			res++
		}
	}

	return res
}
