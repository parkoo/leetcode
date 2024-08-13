package main

// 思路: 二分查找 + 双指针中心扩散

// 时间复杂度: O(logn + k)    空间复杂度: O(1)

func findClosestElements_1(arr []int, k int, x int) []int {
	res := make([]int, 0)

	n := len(arr)
	// 此二分查找返回两个位置l、r，保证[0,l]范围的值小于x, [r,len(arr)-1]范围的值大于等于x
	// 基于这两个位置做选择
	l, r := binarySearch(arr, 0, n-1, x)
	cnt := 0
	for l >= 0 && r < n {
		if abs(arr[l], x) <= abs(arr[r], x) {
			l--
		} else {
			r++
		}
		cnt++
		if cnt == k {
			break
		}
	}

	for cnt < k && l >= 0 {
		l--
		cnt++
	}
	for cnt < k && r < n {
		r++
		cnt++
	}

	res = arr[l+1 : r] // 注意这里是[l+1:r-1], 因为最后l和r在循环中多处理一轮
	return res
}

func abs(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func binarySearch(arr []int, l, r int, t int) (int, int) {
	for l < r {
		mid := l + (r-l)/2

		if arr[mid] >= t {
			r = mid

		} else {
			l = mid + 1
		}
	}

	return l - 1, l
}
