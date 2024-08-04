package main

// 思路: 二分查找

// 时间复杂度：O(lgn)  空间复杂度：O(1)

func search_1(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[0] { // mid 在旋转左半部分, 此时[l, mid] 为有序部分
			if target >= nums[0] && target < nums[mid] { // target 在旋转左半部分，target < nums[mid] (即 target 在 nums[mid] 左边)
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // mid 在旋转右半部分 此时[mid, r] 为有序部分
			if target < nums[0] && target > nums[mid] { // target 在旋转右半部分，target > nums[mid] (即 target 在 nums[mid] 右边)
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
