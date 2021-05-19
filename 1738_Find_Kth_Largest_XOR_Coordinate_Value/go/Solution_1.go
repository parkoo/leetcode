package main

import "math/rand"

// 二维前缀和 + 快排partition
// 异或性质：a^b^b = a, a^0 = a
// 时间复杂度：O(mn), 空间复杂度：O(mn)

func kthLargestValue_1(matrix [][]int, k int) int {
	arr := getXORValues(matrix)
	return getKthNum(arr, 0, len(arr)-1, len(arr)-k)
}

func getXORValues(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	values := make([]int, m*n)
	// 构造二维虚拟边界 a^0 = a
	pre := make([][]int, m+1)
	pre[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		pre[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			pre[i+1][j+1] = pre[i][j+1] ^ pre[i+1][j] ^ pre[i][j] ^ matrix[i][j] // a^b^b = a
			values = append(values, pre[i+1][j+1])
		}
	}

	return values
}

func partition(arr []int, l, r int) int {
	idx := rand.Intn(r-l+1) + l
	arr[l], arr[idx] = arr[idx], arr[l]

	i, j := l+1, r
	for {
		for i <= r && arr[i] < arr[l] {
			i++
		}
		for j >= l+1 && arr[j] > arr[l] {
			j--
		}
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func getKthNum(arr []int, l, r, k int) int {
	p := partition(arr, l, r)
	if p == k {
		return arr[p]
	}
	if p < k {
		return getKthNum(arr, p+1, r, k)
	} else {
		return getKthNum(arr, l, p-1, k)
	}
}
