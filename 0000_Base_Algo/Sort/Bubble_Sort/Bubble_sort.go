package main

import (
	"fmt"
)

// 时间复杂度：O(n^2)  空间复杂度：O(1)
// 稳定排序

func BubbleSort(arr []int) {
	n := len(arr)

	lastSwap := 0
	for i := n - 1; i > 0; i = lastSwap {
		lastSwap = 0
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				lastSwap = j
			}
		}
	}
}

func main() {
	in_1 := []int{7, 0, 4, 6, 2, 4, 9, 1, 3, 2, 5}
	BubbleSort(in_1)
	fmt.Println(in_1)

	in_2 := []int{}
	BubbleSort(in_2)
	fmt.Println(in_2)
}
