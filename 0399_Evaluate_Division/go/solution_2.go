package main

// 思路: BFS

// 时间复杂度：O(ML+Q⋅(L+M)), 其中 M 为边的数量，Q 为询问的数量，L 为字符串的平均长度
// 构建图时，需要处理 M 条边，每条边都涉及到 O(L) 的字符串比较；处理查询时，每次查询首先要进行一次 O(L) 的比较，然后至多遍历 O(M) 条边

// 空间复杂度：O(NL+M)，其中 N 为点的数量，M 为边的数量，L 为字符串的平均长度
// 为了将每个字符串映射到整数，需要开辟空间为 O(NL) 的哈希表；随后，需要花费 O(M) 的空间存储每条边的权重；处理查询时，还需要 O(N) 的空间维护访问队列
// 最终，总的复杂度为 O(NL+M+N)=O(NL+M)

func calcEquation_2(equations [][]string, values []float64, queries [][]string) []float64 {
	res := make([]float64, 0)

	// 字符串映射为节点编号
	toNode := make(map[string]int)

	// 表示带权图的邻接表
	type node struct {
		to     int
		weight float64
	}
	graph := make(map[int][]*node)

	nodeIndex := 0
	for i, equation := range equations {
		x, y := equation[0], equation[1]
		if _, ok := toNode[x]; !ok {
			toNode[x] = nodeIndex
			nodeIndex++
		}
		if _, ok := toNode[y]; !ok {
			toNode[y] = nodeIndex
			nodeIndex++
		}

		nodeY := &node{
			to:     toNode[y],
			weight: values[i],
		}
		graph[toNode[x]] = append(graph[toNode[x]], nodeY)

		nodeX := &node{
			to:     toNode[x],
			weight: 1 / values[i],
		}
		graph[toNode[y]] = append(graph[toNode[y]], nodeX)
	}

	// bfs
	bfs := func(start, end int) float64 {
		// visited[i] = k
		// 若 k == 0, 表示节点i没有被访问过
		// 若 k != 0, 表示节点i已经被访问过，不可重复访问，且从start节点到节点i的总权重为k
		visited := make(map[int]float64)

		q := make([]int, 0)
		q = append(q, start)
		visited[start] = 1

		for len(q) > 0 {
			cur := q[0]
			q = q[1:]

			if cur == end {
				return visited[cur]
			}

			for _, nextNode := range graph[cur] {
				if _, ok := visited[nextNode.to]; !ok {
					q = append(q, nextNode.to)
					visited[nextNode.to] = visited[cur] * nextNode.weight
				}
			}
		}

		return -1.0
	}

	// 查询结果
	for _, query := range queries {
		x, y := query[0], query[1]
		if _, ok := toNode[x]; !ok {
			res = append(res, -1.0)
			continue
		}
		if _, ok := toNode[y]; !ok {
			res = append(res, -1.0)
			continue
		}

		start, end := toNode[x], toNode[y]
		v := bfs(start, end)
		res = append(res, v)
	}

	return res
}
