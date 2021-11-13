package main

import (
	"math/rand"
)

// 时间复杂度：O(nlgn)  空间复杂度：O(n)
// 非稳定排序
// 双路快排

func QuickSort_2(arr []int) {
	quickSort_2(arr, 0, len(arr)-1)
}

func quickSort_2(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition_2(arr, l, r)
	quickSort_2(arr, l, p-1)
	quickSort_2(arr, p+1, r)
}

// 双路快排
func partition_2(arr []int, l, r int) int {
	p := rand.Intn(r-l+1) + l
	arr[l], arr[p] = arr[p], arr[l]

	v := arr[l]
	i, j := l+1, r
	for {
		for i <= r && arr[i] < v {
			i++
		}
		for j >= l+1 && arr[j] > v {
			j--
		}
		// 若只有两个元素且不等式，则初始i=j, 经过上边两个for循环，此时 i>j 直接break
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
