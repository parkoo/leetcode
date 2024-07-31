package main

// 思路:  以board的每个位置作为起点进行全局dfs，在dfs过程中遇到words中的单词则加入结果
// 同时做一些剪枝处理

// 时间复杂度: O(?)    空间复杂度: O(?)

func findWords_2(board [][]byte, words []string) []string {
	res := make([]string, 0)
	if len(board) == 0 || len(board[0]) == 0 {
		return res
	}

	startMap := make(map[byte]struct{}) // 记录每个word的首个字符
	wordMap := make(map[string]bool)    // 记录每个word,以及是否被加入到res中
	for _, word := range words {
		startMap[word[0]] = struct{}{}
		wordMap[word] = false
	}

	m, n := len(board), len(board[0])
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	inArea := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	var dfs func(board [][]byte, x, y int, sub []byte)
	dfs = func(board [][]byte, x, y int, sub []byte) {
		sub = append(sub, board[x][y])
		item := string(sub)

		// word的长度不会超过10，这里做个剪枝
		if len(item) > 10 {
			sub = sub[:len(sub)-1]
			return
		}

		// 加入结果集
		if seen, ok := wordMap[item]; ok && !seen {
			res = append(res, item)
			wordMap[item] = true
		}

		used[x][y] = true
		for k := 0; k < 4; k++ {
			nextX := x + dir[k][0]
			nextY := y + dir[k][1]

			if inArea(nextX, nextY) && !used[nextX][nextY] {
				dfs(board, nextX, nextY, sub)
			}
		}
		used[x][y] = false
		sub = sub[:len(sub)-1]
	}

	// 以board的每个位置作为起点进行dfs
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 剪枝
			if _, ok := startMap[board[i][j]]; !ok {
				continue
			}

			dfs(board, i, j, []byte{})
		}
	}

	return res
}
