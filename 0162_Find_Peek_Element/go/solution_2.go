package main

// 思路: 二分查找
// 关键点：题目说明了可假设 nums[-1] = nums[n] = -∞
// 对于所有有效的 i 都有 nums[i] != nums[i + 1]

// 时间复杂度: O(logn)    空间复杂度: O(1)

func findPeakElement(nums []int) int {

	left, right := 0, len(nums)-1

	for left < right { // 这里保证至少有两个元素参与二分

		// 这里mid一定小于len(nums)-1
		// 在仅有两个元素的情况下，mid取的是左边的元素
		mid := left + (right-left)/2

		// 由于可假设 nums[-1] = nums[n] = -∞， mid处一个可能的解
		if nums[mid] > nums[mid+1] { // mid + 1 不会越界
			right = mid

		} else {
			left = mid + 1
		}
	}
	// 退出时 left == right

	return left
}
