package main

// 双指针 快慢指针
// 时间复杂度：O(n)  空间复杂度：O(1)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 双指针 快慢指针
	// 保证nums[0,i]中为不重复的严格递增的数字
	i, j := 0, 0
	for ; j < len(nums); j++ {
		if nums[j] > nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
