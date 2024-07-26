package main

// 思路: 环形区间 说明说明数组首尾不能同时被偷取
// 分别对[0, len(nums)-2] 和 [1, len(nums)-1] 进行偷取，取最大值即可

// 时间复杂度: O(n)    空间复杂度: O(1)

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	return max(robWithRange(nums, 0, n-2), robWithRange(nums, 1, n-1))
}

func robWithRange(nums []int, start, end int) int {
	pre1, pre2 := 0, 0
	res := 0
	for i := start; i <= end; i++ {
		res = max(pre1+nums[i], pre2)

		pre1 = pre2
		pre2 = res
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
