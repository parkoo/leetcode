package main

import (
	"fmt"
)

func main() {
	in_1 := []int{7, 0, 4, 6, 2, 4, 9, 1, 3, 2, 5}
	QuickSort_1(in_1)
	fmt.Println(in_1)

	in_2 := []int{}
	QuickSort_1(in_2)
	fmt.Println(in_2)
}
