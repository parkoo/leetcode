package main

import "math"

// 思路：看组时数组的组合问题 可重复选取元素的组合问题同lc39  但时会超时

func coinChange(coins []int, amount int) int {
	res := math.MaxInt

	var backtrack func(nums []int, sub []int, sum int, start int)
	backtrack = func(nums []int, sub []int, sum int, start int) {
		if sum > amount {
			return
		}
		if sum == amount {
			item := make([]int, 0)
			item = append(item, sub...)
			curLen := len(item)
			if curLen < res {
				res = curLen
			}
			return
		}

		for i := start; i < len(nums); i++ {
			sub = append(sub, nums[i])
			sum += nums[i]
			// start传入i, 保证下层节点与当前节点从同一个位置开始，避免结果重复 (如：[2,2,3]、[2,3,2]、[3,2,2])
			// 组合问题中 在回溯函数中需要传入起始位置，若数组中的元素可以重复使用，则向下层函数传入当前位置i,
			// 若数组中的元素可以重复使用，则向下层函数传入当前位置i的下一个位置，即i+1
			backtrack(nums, sub, sum, i)
			sub = sub[:len(sub)-1]
			sum -= nums[i]
		}
	}

	backtrack(coins, []int{}, 0, 0)

	if res == math.MaxInt {
		res = -1
	}
	return res
}
