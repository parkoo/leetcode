package main

// 数组排序
// 迭代实现 自底向上排序
// 稳定排序
// 时间复杂度；O(nlgn)  空间复杂度：O(n)

func MergeSort_2(arr []int) {
	merger_2(arr)
}

func merger_2(arr []int) {
	n := len(arr)
	// 对 arr[i, i+step-1]  arr[i+step, i+srep+step-1] 归并
	// step<n, 而非 step<=(n+1)/2, 因为step是翻倍增长的，
	// 若n=5, 最后要对arr[0, 3] arr[4]归并，此时step = 4 , 若step<=(n+1)/2， 则step无法取到4
	for step := 1; step < n; step = step * 2 {
		for i := 0; i < n-step; i = i + step + step {
			l, r := i, min(i+step+step-1, n-1)
			m := i + step - 1
			merge_2(arr, l, r, m)
		}
	}
}

func merge_2(arr []int, l, r, m int) {
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = arr[i]
	}

	i, j := l, m+1
	for k := l; k <= r; k++ {
		if i > m {
			arr[k] = temp[j-l]
			j++
		} else if j > r {
			arr[k] = temp[i-l]
			i++
		} else {
			if temp[i-l] < temp[j-l] {
				arr[k] = temp[i-l]
				i++
			} else {
				arr[k] = temp[j-l]
				j++
			}
		}
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
