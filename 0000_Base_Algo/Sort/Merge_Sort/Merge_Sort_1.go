package main

// 数组排序
// 递归实现 自顶向下排序
// 稳定排序
// 时间复杂度；O(nlgn)  空间复杂度：O(n)

func MergeSort_1(arr []int) {
	merger_1(arr, 0, len(arr)-1)
}

func merger_1(arr []int, l, r int) {
	if l >= r {
		return
	}

	m := l + (r-l)/2
	merger_1(arr, l, m)
	merger_1(arr, m+1, r)
	merge_1(arr, l, r, m)
}

func merge_1(arr []int, l, r, m int) {
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
			if temp[i-l] > temp[j-l] {
				arr[k] = temp[j-l]
				j++
			} else {
				arr[k] = temp[i-l]
				i++
			}
		}
	}
}
