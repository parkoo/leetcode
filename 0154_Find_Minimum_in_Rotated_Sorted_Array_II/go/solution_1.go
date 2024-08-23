package main

// 思路:  二分查找

// 时间复杂度: O(n)(最坏)    空间复杂度: O(1)

func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	left, right := 0, n-1
	for left < right {
		flag := nums[right] // 以nums[right]作为标定点

		mid := left + (right-left)/2

		// 由于存在重复元素，可能出现 arr[left] == arr[right] == arr[mid] 的情况，此时可以将右端点左移一位
		// 处理重复元素
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			right--
			continue
		}

		if nums[mid] > flag {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
