package main

import (
	"math/rand"
)

// 时间复杂度：O(nlgn)  空间复杂度：O(lgn)
// 非稳定排序
// 三路快排

func QuickSort_3(arr []int) {
	quickSort_3(arr, 0, len(arr)-1)
}

func quickSort_3(arr []int, l, r int) {
	if l >= r {
		return
	}

	// partition过程 需返回两个位置点i、j, 最终使得arr[l,i-1]<v && arr[i,j-1]==v && arr[j,r]>v
	// 随机选取标定点 防止递归树不平衡
	p := rand.Intn(r-l+1) + l
	arr[p], arr[l] = arr[l], arr[p]

	v := arr[l]
	i, j, k := l, r+1, l+1
	for k < j {
		if arr[k] < v {
			i++
			arr[i], arr[k] = arr[k], arr[i]
			k++
		} else if arr[k] > v {
			j--
			arr[k], arr[j] = arr[j], arr[k]
		} else {
			k++
		}
	}
	arr[l], arr[i] = arr[i], arr[l] // 交换后 arr[i]==v

	quickSort_3(arr, l, i-1)
	quickSort_3(arr, j, r)
}
