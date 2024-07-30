package main

// 思路:  最长非严格递增子序列
// 思路类似 lc0300最长上升子序列

// 时间复杂度: O(n^2)    空间复杂度: O(n)

func longestObstacleCourseAtEachPosition_1(obstacles []int) []int {
	n := len(obstacles)

	res := make([]int, n)
	res[0] = 1

	for i := 1; i < n; i++ {
		res[i] = 1
		for j := 0; j < i; j++ {
			if obstacles[i] >= obstacles[j] {
				if res[j]+1 > res[i] {
					res[i] = res[j] + 1
				}

			}
		}
	}

	return res
}
