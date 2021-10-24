package main

// 双指针 快慢指针
// 时间复杂度：O(n)  空间复杂度：O(1)

func moveZeroes_1(nums []int) {
	// 双指针 快慢指针
	// 保证nums[0,i]区间为非零数字 初始 i = -1
	i, j := -1, 0
	for ; j < len(nums); j++ {
		if nums[j] != 0 {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
