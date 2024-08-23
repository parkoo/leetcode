package main

// 思路:  二分查找

// 时间复杂度: O(n)(最坏)    空间复杂度: O(1)

func search(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}

	left, right := 0, n-1
	for left < right {
		// 注意这里的flag是arr[left] 而不是arr[0]
		flag := nums[left]

		mid := left + (right-left)/2
		// 由于存在重复元素，可能出现 arr[left] == arr[right] == arr[mid] 的情况，此时可以将右端点左移一位
		// 处理重复元素
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			right--
			continue
		}

		if nums[mid] >= flag {
			if target >= flag && target <= nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}

		} else if nums[mid] < flag {
			if target < flag && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid
			}

		}
	}

	return nums[left] == target
}
