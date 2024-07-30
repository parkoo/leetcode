package main

// 思路: 最长非严格递增子序列
// 思路类似 lc0300最长上升子序列
// 注意两个不同的地方

// 时间复杂度: O(nlgn)    空间复杂度: O(n)

func longestObstacleCourseAtEachPosition(obstacles []int) []int {

	n := len(obstacles)
	res := make([]int, n)

	d := make([]int, n+1)
	maxLen := 1
	d[maxLen] = obstacles[0]
	res[0] = 1

	for i := 1; i < n; i++ {
		cur := obstacles[i]
		if cur >= d[maxLen] { // 注意1:找非严格递增子序列时，这里为 >=
			maxLen++
			d[maxLen] = cur
			res[i] = maxLen

		} else {
			// 注意2:找非严格递增子序列时，cur+1
			// 因为可能出现多个length对应一个cur, 如 [2, 3, 3, 1] -> d[1] = 2, d[2] = 3, d[3] = 3
			// 使用 cur + 1 可以保证得到的位置一定是比cur大的第一个数（不等于cur）
			index := binarySearch(d, maxLen, cur+1)
			d[index] = cur
			res[i] = index
		}
	}

	return res
}

func binarySearch(arr []int, maxLen int, target int) int {
	left, right := 1, maxLen

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] >= target {
			right = mid

		} else {
			left = mid + 1
		}
	}

	return left
}
