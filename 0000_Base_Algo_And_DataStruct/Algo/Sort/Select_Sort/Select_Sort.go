package main

import (
	"fmt"
)

// 时间复杂度：O(n^2)  空间复杂度：O(1)
// 非稳定排序, 例: 5 8 5 2 9

func SelectSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	for i := 0; i < len(arr); i++ {
		curIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[curIndex] {
				curIndex = j
			}
		}
		arr[curIndex], arr[i] = arr[i], arr[curIndex]
	}
}

func main() {
	in_1 := []int{7, 0, 4, 6, 2, 4, 9, 1, 3, 2, 5}
	SelectSort(in_1)
	fmt.Println(in_1)

	in_2 := []int{}
	SelectSort(in_2)
	fmt.Println(in_2)
}
