package main

import "math"

// 思路: BFS

// 时间复杂度: O(n)    空间复杂度: O(n)

func minJumps_1(arr []int) int {
	n := len(arr)

	indexMap := make(map[int][]int) // 记录等值元素的所有下标
	dist := make([]int, n)          // 记录到达某个位置的最小步数
	visited := make(map[int]bool)   // 记录是否访问过

	for i, num := range arr {
		indexMap[num] = append(indexMap[num], i)
		dist[i] = math.MaxInt // 初始化为最大值
	}

	// BFS队列
	dq := make([]int, 0)
	dq = append(dq, 0)
	dist[0] = 0
	visited[0] = true

	for len(dq) > 0 {
		cur := dq[0]
		dq = dq[1:]

		curStep := dist[cur]
		// 到达数组最后一个位置
		if cur == n-1 {
			return curStep
		}

		// 获取后继节点 nextList
		nextList := indexMap[arr[cur]]
		// 注意访问后要做删除，避免重复访问等值位置，降低时间复杂度
		// 删除这一步是降低时间复杂度的关键 O(n^2) > O(n)
		delete(indexMap, arr[cur])
		nextList = append(nextList, cur+1)
		if cur > 0 {
			nextList = append(nextList, cur-1)
		}

		for _, next := range nextList {
			if visited[next] {
				continue
			}

			dq = append(dq, next)
			dist[next] = curStep + 1
			visited[next] = true
		}
	}

	return -1
}
