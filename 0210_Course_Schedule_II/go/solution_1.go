package main

// 思路: 图的拓扑排序，bfs

// 时间复杂度: O(m+n)    空间复杂度: O(m+n)

func findOrder_1(numCourses int, prerequisites [][]int) []int {
	res := make([]int, 0)
	n := numCourses

	indegrees := make([]int, n)  // 入度表
	graph := make(map[int][]int) // 图的邻接表
	for _, prerequisite := range prerequisites {
		// cur -> pre, cur 依赖 pre
		cur, pre := prerequisite[0], prerequisite[1]
		indegrees[pre]++
		graph[cur] = append(graph[cur], pre)
	}

	// bfs
	bfs := func() {
		q := make([]int, 0)
		for node, degree := range indegrees {
			if degree == 0 {
				q = append(q, node)
			}
		}

		for len(q) > 0 {
			cnt := len(q)
			for i := 0; i < cnt; i++ {
				curNode := q[0]
				q = q[1:]
				res = append(res, curNode)

				for _, nextNode := range graph[curNode] {
					indegrees[nextNode]--
					if indegrees[nextNode] == 0 {
						q = append(q, nextNode)
					}
				}
			}
		}
	}

	bfs()

	// 判断是否存在有效结果
	if len(res) == n {
		reverse(res)
		return res
	}

	return []int{}
}

func reverse(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
