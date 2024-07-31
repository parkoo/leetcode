package main

// 思路: BFS
// 类似 lc0127

// 时间复杂度: O(mn)    空间复杂度: O(mn)

func minMutation_1(startGene string, endGene string, bank []string) int {

	// 判断基因a是否经过一次转换可以变成基因b
	canMutate := func(a, b string) bool {
		diffCnt := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				diffCnt++
			}
		}

		return diffCnt == 1
	}

	// bfs
	visited := make(map[string]bool)
	q := make([]string, 0)
	q = append(q, startGene)

	step := 0
	for len(q) > 0 {
		cnt := len(q)

		for i := 0; i < cnt; i++ {
			cur := q[0]
			q = q[1:]
			if cur == endGene {
				return step
			}

			for _, next := range bank {
				if canMutate(cur, next) && !visited[next] {
					q = append(q, next)
					visited[next] = true
				}
			}
		}
		step++
	}

	return -1
}
