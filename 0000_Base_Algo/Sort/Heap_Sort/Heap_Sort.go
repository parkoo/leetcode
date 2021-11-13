package main

import (
	"fmt"
)

// 时间复杂度：O(nlgn)  空间复杂度：O(1)
// 非稳定排序

func HeapSort(arr []int) {
	heapify(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		shiftDown(arr, 0, i)
	}
}

// 一个大小为n的堆，其第一个非叶子节点的索引为n/2
// heapify 就是由下向上 由右至左 对非叶子节点进行shiftDown
func heapify(arr []int) {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		shiftDown(arr, i, n)
	}
}

// 对arr[0, n)中的第k个元素进行shiftDown
func shiftDown(arr []int, k, n int) {
	for 2*k+1 < n {
		j := 2*k + 1
		if j+1 < n && arr[j+1] > arr[j] {
			j += 1
		}
		if arr[k] > arr[j] {
			break
		}

		arr[j], arr[k] = arr[k], arr[j]
		k = j
	}
}

func main() {
	in_1 := []int{7, 0, 4, 6, 2, 4, 9, 1, 3, 2, 5}
	HeapSort(in_1)
	fmt.Println(in_1)

	in_2 := []int{}
	HeapSort(in_2)
	fmt.Println(in_2)
}
