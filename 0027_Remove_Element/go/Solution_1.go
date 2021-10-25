package main

// 双指针 快慢指针
// 时间复杂度：O(n)  空间复杂度；O(1)

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	// 双指针 快慢指针
	// 保证nums[0,i]区间内是不为val的元素
	i, j := -1, 0
	for ; j < len(nums); j++ {
		if nums[j] != val {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return i + 1
}
