package main

// 思路: 最终的结果在环形区间可以分为两种情况，一种是中间的正常区间，一种是头尾区间
// 对正常区间的最小值取反即可得到头尾区间的最大值
// 分别计算正常区间的最大值和最小值
// 注意数组中均为负数的情况，此时对正常区间的最小值取反的值为0, 需要特殊处理

// 时间复杂度: O(n)    空间复杂度: O(1)

func maxSubarraySumCircular_1(nums []int) int {
	res := 0
	if len(nums) == 0 {
		return res
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	preMin, preMax := nums[0], nums[0]
	minRes, maxRes := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		curMin := min(preMin+nums[i], nums[i])
		curMax := max(preMax+nums[i], nums[i])

		preMin, preMax = curMin, curMax

		minRes = min(minRes, curMin)
		maxRes = max(maxRes, curMax)
	}

	res = max(sum-minRes, maxRes)
	// 特殊情况，nums中均是负数
	// 此时maxRes为所有负数中最大的一个
	if maxRes < 0 {
		res = maxRes
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
