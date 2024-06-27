package main

// 思路：动态规划  注意和 “lc53求和最大的子数组” 比较，求和不会受符号变化的影响
// 注意，因为存在负数，当前位置的最大值可以由前一个位置的最大值获得，也可能由前一个位置的最小值获得
// 所以，对于遍历到的每个位置，我们都需要记录下“以当前位置为结尾的子数组的最大乘积值” 和 “以当前位置为结尾的子数组的最小乘积值”
// 定义状态dpMax[i]表示: 以当前位置i为结尾的子数组的最大乘积值, 定义状态dpMin[i]表示: 以当前位置i为结尾的子数组的最小乘积值
// 则状态转移方程为: dpMax[i] = max(nums[i], nums[i]*dpMax[i-1],  nums[i]*dpMin[i-1]), dpMin[i] = min(nums[i], nums[i]*dpMax[i-1],  nums[i]*dpMin[i-1])
// 由于计算当前位置的状态只需要依赖前一个位置的状态，可以进一步优化空间复杂度

// 时间复杂度: O(n)    空间复杂度: O(1)

func maxProduct_1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 一直相乘避免溢出转换成float64
	res := float64(nums[0])
	preMax, preMin := float64(nums[0]), float64(nums[0])
	for i := 1; i < n; i++ {
		curMax := max(max(float64(nums[i]), float64(nums[i])*preMax), float64(nums[i])*preMin) // 以当前位置为结尾的子数组的最大乘积
		curMin := min(min(float64(nums[i]), float64(nums[i])*preMax), float64(nums[i])*preMin) // 以当前位置为结尾的子数组的最小乘积

		preMax, preMin = curMax, curMin

		if curMax > res {
			res = curMax
		}
	}

	return int(res)
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
