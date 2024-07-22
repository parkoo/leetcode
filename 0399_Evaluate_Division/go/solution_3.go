package main

// 思路: DFS

// 时间复杂度: O(?)    空间复杂度: O(?)

func calcEquation_3(equations [][]string, values []float64, queries [][]string) []float64 {
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

	// 记录dfs时是否访问过该节点，防止在环里陷入死循环
	visited := make(map[int]bool)

	// dfs
	var dfs func(start, end int, accWeight float64) float64
	dfs = func(start, end int, accWeight float64) float64 {
		if start == end {
			return accWeight
		}

		for _, nextNode := range graph[start] {
			if !visited[nextNode.to] {
				visited[start] = true
				subRes := dfs(nextNode.to, end, accWeight*nextNode.weight)
				visited[start] = false
				if subRes != -1.0 {
					return subRes
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

		v := dfs(toNode[x], toNode[y], 1.0)
		res = append(res, v)
	}

	return res
}
