package main

import (
	//	"fmt"
	"math/rand"
)

// 时间复杂度：O(nlgn)  空间复杂度：O(lgn)
// 非稳定排序
// 单路快排

func QuickSort_1(arr []int) {
	quickSort_1(arr, 0, len(arr)-1)
}

func quickSort_1(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition_1(arr, l, r)
	quickSort_1(arr, l, p-1)
	quickSort_1(arr, p+1, r)
}

// 单路快排
func partition_1(arr []int, l, r int) int {
	p := rand.Intn(r-l+1) + l
	arr[l], arr[p] = arr[p], arr[l]

	v := arr[l]
	i, j := l, l+1
	for ; j <= r; j++ {
		if arr[j] < v {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[l], arr[i] = arr[i], arr[l]

	return i
}

// func main() {
// 	in_1 := []int{7, 0, 4, 6, 2, 4, 9, 1, 3, 2, 5}
// 	QuickSort_1(in_1)
// 	fmt.Println(in_1)

// 	in_2 := []int{}
// 	QuickSort_1(in_2)
// 	fmt.Println(in_2)
// }
