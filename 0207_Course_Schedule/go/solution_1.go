package main

// 思路:

// 时间复杂度: O(?)    空间复杂度: O(?)

func canFinish_1(numCourses int, prerequisites [][]int) bool {

	if len(prerequisites) == 0 { // 课程之间无依赖关系，可直接学习
		return true
	}

	// 构建入度数组、临接表
	indegrees := make([]int, numCourses)
	adjacency := make(map[int][]int)
	for _, item := range prerequisites {
		indegrees[item[0]]++
		adjacency[item[1]] = append(adjacency[item[1]], item[0])
	}

	// 设置 bfs 队列, 初始节点入队
	q := make([]int, 0)
	for node, indegree := range indegrees {
		if indegree == 0 { // 入度为0, 不依赖其他节点，可入队
			q = append(q, node)
		}
	}

	cnt := 0 // 记录遍历节点个数
	for len(q) > 0 {
		// 节点出队
		cur := q[0]
		q = q[1:]
		cnt++

		// 更新下游节点next的入度, 下游节点入度减1, 若减为0则入队
		for _, next := range adjacency[cur] {
			indegrees[next]--
			if indegrees[next] == 0 {
				q = append(q, next)
			}
		}
	}

	return cnt == numCourses
}
