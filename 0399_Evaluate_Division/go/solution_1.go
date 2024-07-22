package main

// 思路: 带权并查集

// 时间复杂度：O((N+Q)logA)
// 构建并查集 O(NlogA) ，这里 N 为输入方程 equations 的长度，每一次执行合并操作的时间复杂度是 O(logA)，这里 A 是 equations 里不同字符的个数
// 查询并查集 O(QlogA)，这里 Q 为查询数组 queries 的长度，每一次查询时执行「路径压缩」的时间复杂度是 O(logA)

// 空间复杂度：O(A)
// 创建字符与 nodeIndex 的对应关系 toNode 长度为 A，并查集底层使用的两个数组 parent 和 weight 存储每个变量的连通分量信息，parent 和 weight 的长度均为 A

func calcEquation_1(equations [][]string, values []float64, queries [][]string) []float64 {
	res := make([]float64, 0)

	toNode := make(map[string]int)
	nodeIndex := 0
	// 初始化并查集
	unionFind := NewUnionFind(2 * len(equations))

	// 预处理
	for i, equation := range equations {
		if _, ok := toNode[equation[0]]; !ok {
			toNode[equation[0]] = nodeIndex
			nodeIndex++
		}

		if _, ok := toNode[equation[1]]; !ok {
			toNode[equation[1]] = nodeIndex
			nodeIndex++
		}

		unionFind.Union(toNode[equation[0]], toNode[equation[1]], values[i])
	}

	// 查询结果
	for _, query := range queries {
		nodeA, ok := toNode[query[0]]
		if !ok {
			res = append(res, -1.0)
			continue
		}

		nodeB, ok := toNode[query[1]]
		if !ok {
			res = append(res, -1.0)
			continue
		}

		v := unionFind.IsConnected(nodeA, nodeB)
		res = append(res, v)
	}

	return res
}

// 带权重的并查集
type UnionFind struct {
	parent []int
	weight []float64
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	weight := make([]float64, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		weight[i] = 1.0
	}

	return &UnionFind{
		parent: parent,
		weight: weight,
	}
}

// 递归压缩的find
func (u *UnionFind) find(node int) int {
	if u.parent[node] != node {
		nextParent := u.find(u.parent[node])
		// 这里需要先赋值weight、后赋值parent
		u.weight[node] *= u.weight[u.parent[node]]
		u.parent[node] = nextParent
	}

	return u.parent[node]
}

// 返回节点之间的权重，若节点不在一个集合则返回 -1.0
func (u *UnionFind) IsConnected(nodeA, nodeB int) float64 {
	rootA := u.find(nodeA)
	rootB := u.find(nodeB)
	if rootA != rootB {
		return -1.0
	}

	return u.weight[nodeA] / u.weight[nodeB]
}

// 带权重的合并
func (u *UnionFind) Union(nodeA, nodeB int, weight float64) {
	rootA := u.find(nodeA)
	rootB := u.find(nodeB)
	if rootA == rootB {
		return
	}

	u.parent[rootA] = rootB
	u.weight[rootA] = u.weight[nodeB] * weight / u.weight[nodeA]
}
