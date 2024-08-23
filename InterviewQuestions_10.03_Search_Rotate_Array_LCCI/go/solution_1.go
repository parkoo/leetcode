package main

// 思路:  二分查找

// 时间复杂度: O(n)(最坏)    空间复杂度: O(1)

func search(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	left, right := 0, n-1
	for left < right {
		// 注意这里的flag是arr[left] 而不是arr[0]
		flag := arr[left]

		mid := left + (right-left)/2

		// 由于存在重复元素，可能出现 arr[left] == arr[right] == arr[mid] 的情况，此时可以将右端点左移一位
		// 处理重复元素
		if arr[left] == arr[right] && arr[mid] == arr[right] {
			right--
			continue
		}

		if arr[mid] >= flag {
			if target >= flag && target <= arr[mid] {
				right = mid
			} else {
				left = mid + 1
			}

		} else {
			if target < flag && target > arr[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	if arr[left] == target {
		return left
	}

	return -1
}
