package main

// 思路: 二分法寻找最优区间的左边界

// 时间复杂度: O(logn)    空间复杂度: O(1)

func findClosestElements_2(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-k
	for l < r {
		mid := l + (r-l)/2

		// 这里是关键，要确定区间[l:l+k]是需要左移还是右移
		// 理想最优结果是 x - arr[mid] == arr[mid+k] - x
		// 当不存在理想最优结果时，需要满足 x - arr[mid] < arr[mid+k] - x
		// 当 x - arr[mid] > arr[mid+k] - x, 区间一定需要右移
		if x-arr[mid] > arr[mid+k]-x {
			l = mid + 1
		} else {
			r = mid
		}
	}

	res := arr[l : l+k]
	return res
}
