package main

// 思路: BFS
// 将棋盘抽象成一个包含 N^2个节点的有向图，对于每个节点 x，若 x+i(1≤i≤6) 上没有蛇或梯子，则连一条从 x 到 x+i 的有向边
// 否则记蛇梯的目的地为 y，连一条从 x 到 y 的有向边
// 原问题等价于在有向图上求出从 1 到 N^2的最短路长度

// 时间复杂度: O(n^2)    空间复杂度: O(n^2)

func snakesAndLadders_1(board [][]int) int {
	n := len(board)

	// 将id转换为坐标
	id2rc := func(id, n int) (int, int) {
		r := (id - 1) / n
		c := (id - 1) % n
		if r%2 == 1 {
			c = n - c - 1
		}
		r = n - r - 1

		return r, c
	}

	// 仅记录到达过的节点，中间的梯子或蛇的中转节点不计入
	// 因为当前到达的是经过梯子或蛇中转后的节点，而后续可能还会使用该中转节点
	//（可以经过一次中转后再到达该中转节点， 则该中转节点不可在中转，可以当做普通节点）
	// [-1,1,1,1]  [16,15,14,13]
	// [-1,7,1,1]  [09,10,11,12]
	// [16,1,1,1]  [08,07,06,05]
	// [-1,1,9,1]  [01,02,03,04]
	// 01 >> 09(经过03中转) >> 07(经过10中转) >> 16(经过08中转)
	// 第二步到达的07本来是个中转节点，但由于到达07是经过10中转过的，而一次跳跃只能中转一次，所以07不在做中转，可以在07处停留
	// 也说明中转过程中的中转节点先不需要标记为已经访问
	visited := make(map[int]bool)
	q := make([]int, 0)
	q = append(q, 1)
	visited[1] = true

	step := 0
	for len(q) > 0 {
		cnt := len(q)
		for i := 0; i < cnt; i++ {
			curNodeId := q[0]
			q = q[1:]
			if curNodeId == n*n {
				return step
			}

			for k := 1; k <= 6; k++ {
				nextNodeId := curNodeId + k
				if nextNodeId > n*n {
					break
				}

				// 处理有蛇或梯子的情况
				// 若存在蛇或梯子，必须经过蛇或者梯子
				// 蛇或梯子在一轮选择中只能使用一次
				nextX, nextY := id2rc(nextNodeId, n)
				if board[nextX][nextY] != -1 {
					nextNodeId = board[nextX][nextY]
				}

				if !visited[nextNodeId] {
					q = append(q, nextNodeId)
					visited[nextNodeId] = true
				}
			}
		}
		if len(q) > 0 {
			step++
		}
	}

	return -1
}
