package main

// 思路: 动态规划
// 对 slolution_1 优化空间维度

// 时间复杂度: O(n)    空间复杂度: O(1)

func numberOfArithmeticSlices_2(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	res := 0
	pre, cur := 0, 0
	for i := 2; i < n; i++ {
		cur = 0 // 这里需要设置初值
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			cur = pre + 1
		}
		pre = cur

		res += cur
	}

	return res
}
