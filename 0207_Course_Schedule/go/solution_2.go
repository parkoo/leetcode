package main

// 思路: DFS 判断图中是否存在环，求图的拓扑排列

// 时间复杂度: O(v+e)    空间复杂度: O(v+e)  其中，v为课程数(图的节点数)，e为课程之间的关系数(图的边数)

func canFinish_2(numCourses int, prerequisites [][]int) bool {
	valid := true

	// 节点的状态数组, 0:“未开始访问” 1:”访问中“ 2:”完成访问“, 初始化均为0
	visited := make([]int, numCourses)

	// 构建邻接表
	adjacency := make(map[int][]int)
	for _, item := range prerequisites {
		adjacency[item[1]] = append(adjacency[item[1]], item[0])
	}

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = 1 // 设置为”访问中“

		for _, next := range adjacency[node] {
			if visited[next] == 0 { // 该下游节点未开始访问，则开始对该节点进行访问
				dfs(next)
				if !valid { // 访问下游节点后，判断有效性
					return
				}

			} else if visited[next] == 1 { // 该下游节点同时正在被访问中，表示该图存在环，不存在拓扑排序，直接返回
				valid = false
				return
			}

			// 下游节点已经完成访问，则无需操作
		}

		// 当前节点的所有下游节点均顺利访问成功，表示当前节点访问成功完成
		visited[node] = 2
	}

	// 对未开始的节点做遍历访问
	for node := 0; node < numCourses; node++ {
		if visited[node] == 0 && valid {
			dfs(node)
		}
	}

	return valid
}
