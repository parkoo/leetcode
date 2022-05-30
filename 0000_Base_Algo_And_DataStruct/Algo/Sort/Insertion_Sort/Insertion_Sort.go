package main

import (
	"fmt"
)

// 时间复杂度：O(n^2)  空间复杂度：O(1)
// 稳定排序

func InsertionSort(arr []int) {
	if len(arr) == 0 || len(arr) == 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		j := i      // j为arr[i]插入的位置
		v := arr[j] // arr[i]会被覆盖 需记录
		for ; j > 0 && arr[j-1] > v; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = v
	}
}

func main() {
	in_1 := []int{7, 0, 4, 6, 2, 4, 9, 1, 3, 2, 5}
	InsertionSort(in_1)
	fmt.Println(in_1)

	in_2 := []int{}
	InsertionSort(in_2)
	fmt.Println(in_2)
}
