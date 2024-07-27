package main

// 思路:  子序列问题  当前节点考虑

// 时间复杂度: O(n^2)    空间复杂度: O(n)

func wiggleMaxLength_1(nums []int) int {
	res := 0
	n := len(nums)
	if n == 0 {
		return res
	}

	// dp0[i] 表示以nums[i]为结尾(将nums[i]选入摆动序列中)的且nums[i]为下降点的最长摆动序列的长度为dp0[i]
	// dp1[i] 表示以nums[i]为结尾(将nums[i]选入摆动序列中)的且nums[i]为上升点的最长摆动序列的长度为dp1[i]
	dp0 := make([]int, len(nums))
	dp1 := make([]int, len(nums))
	dp0[0], dp1[0] = 1, 1
	res = 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp1[i] = max(dp1[i], dp0[j]+1)
			}
			if nums[i] < nums[j] {
				dp0[i] = max(dp0[i], dp1[j]+1)
			}
		}

		res = max(res, max(dp0[i], dp1[i]))
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
