package main

// 思路: 数组中的值在[1, n]之间，n为数组长度，数组中元素归位
// 此方法可以解决问题，但修改了原数组，不符合题目要求

// 时间复杂度: O(n)    空间复杂度: O(1)

func findDuplicate_1(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return nums[i]
		}
	}

	return 0
}
