package main

// 思路: 计算过程中周围细胞状态相互影响，最直接简单的方法就是复制一份初始态，通过复制的初始态计算下一轮状态，但这样空间复杂度为O(m*n)
// 考虑将计算到的下一轮状态额外存储，即计算过程中同时记录“当前态”和“下一轮状态”（混合态），细胞状态的“下一轮状态”通过周围细胞的“当前态”计算
// 经过遍历一次计算完成，再遍历一次将混合态更新为下一轮状态
// 这里，将下一轮状态记录在十位数上，个位数不变依然为当前轮状态

// 时间复杂度: O(m*n)    空间复杂度: O(1)

func gameOfLife(board [][]int) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])

	// 校验坐标[x,y]是否越界
	inArea := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	// 计算[x,y]处细胞的混合状态（当前状态、下一轮状态共存态）
	// 将下一轮状态记录在十位数上，个位数不变依然为当前轮状态
	// 这样即得到了下一轮状态，又可以根据个位上的数字（当前轮状态）继续对其他细胞作计算
	getMixedStatus := func(x, y int) {
		liveCnt := 0 // 统计初始状态下，当前细胞周围活细胞的个数

		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				if !inArea(i, j) {
					continue
				}
				if i == x && j == y {
					continue
				}
				if board[i][j]%10 == 1 {
					liveCnt++
				}
			}
		}

		// 根据‘当前细胞状态’ 和 ‘当前细胞周围活细胞的个数’ 计算当前细胞的下一轮状态
		// 将下一轮状态记录在十位数上，避免影响当前轮下一个细胞得计算
		curStatus := board[x][y] % 10
		nextStatus := 0
		if (liveCnt == 2 || liveCnt == 3) && curStatus == 1 {
			nextStatus = 1
		}
		if liveCnt == 3 && curStatus == 0 {
			nextStatus = 1
		}

		// 更新混合状态 (十位数字为下一轮状态, 个位数字为当前状态)
		board[x][y] = nextStatus*10 + curStatus
	}

	// 计算每个坐标的混合状态
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			getMixedStatus(i, j)
		}
	}

	// 将每个坐标的混合状态更新为下一轮状态（取十位数上的数字）
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] /= 10
		}
	}
}
