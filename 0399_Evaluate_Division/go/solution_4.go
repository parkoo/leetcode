package main

// 思路: Floyd算法
// https://www.cnblogs.com/wangyuliang/p/9216365.html

// 时间复杂度：O(ML+N^3+QL)
// 构建图需要 O(ML) 的时间, Floyd 算法需要 O(N^3) 的时间, 处理查询时，单次查询只需要 O(L) 的字符串比较以及常数时间的额外操作

// 空间复杂度：O(NL+N^2)

func calcEquation_4(equations [][]string, values []float64, queries [][]string) []float64 {
	res := make([]float64, 0)

	// 字符串映射为节点编号
	toNode := make(map[string]int)
	nodeIndex := 0
	for _, equation := range equations {
		x, y := equation[0], equation[1]
		if _, ok := toNode[x]; !ok {
			toNode[x] = nodeIndex
			nodeIndex++
		}
		if _, ok := toNode[y]; !ok {
			toNode[y] = nodeIndex
			nodeIndex++
		}
	}

	// 建图  graph[i][j] = k, 表示节点i指向节点j的权重值为k
	graph := make([][]float64, len(toNode))
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]float64, len(toNode))
	}
	for i, equation := range equations {
		x, y := equation[0], equation[1]
		graph[toNode[x]][toNode[y]] = values[i]
		graph[toNode[y]][toNode[x]] = 1 / values[i]
	}

	// floyd 算法  计算节点i到节点j的路径
	// 从节点i到达节点j，中间需经过节点k中转，注意这里的顺序
	for k := 0; k < len(graph); k++ {
		for i := 0; i < len(graph); i++ {
			for j := 0; j < len(graph); j++ {
				if graph[i][k] > 1e-6 && graph[k][j] > 1e-6 { //  1/values[i] 会丢失精度，需要做精度处理
					graph[i][j] = graph[i][k] * graph[k][j]
				}
			}
		}
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

		v := graph[toNode[x]][toNode[y]]
		if v == 0 {
			v = -1.0
		}
		res = append(res, v)
	}

	return res
}
