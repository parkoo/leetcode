package main

// 思路: 数组归位法
// 对于一个长度为 N 的数组，其中没有出现的最小正整数只能在 [1,N+1]中。
// 这是因为如果 [1,N] 都出现了，那么答案是 N+1，否则答案是 [1,N] 中没有出现的最小正整数

// 时间复杂度: O(n) 对每个元素作归位操作   空间复杂度: O(1)

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		// 注意条件 “nums[nums[i]-1] != nums[i]”, 确保当前元素 nums[i] 待交换的位置 nums[i]-1 上的数 nums[nums[i]-1] 不等于 nums[i], 否则交换会死循环
		// 即判断待交换的位置上的数是否是正确的，如果是正确的则不需要交换
		for (nums[i] >= 1 && nums[i] <= n) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}
