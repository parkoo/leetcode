package main

// 思路:  记录上升长度和下降长度

// 时间复杂度: O(n)    空间复杂度: O(1)

func longestMountain_2(arr []int) int {
	res := 0

	n := len(arr)
	i := 1
	for i < n {
		inc, dec := 0, 0

		for i < n && arr[i] > arr[i-1] {
			inc++
			i++
		}
		for i < n && arr[i] < arr[i-1] {
			dec++
			i++
		}

		if inc > 0 && dec > 0 {
			curLen := inc + dec + 1
			if curLen > res {
				res = curLen
			}
		}

		for i < n && arr[i] == arr[i-1] {
			i++
		}
	}

	return res
}
