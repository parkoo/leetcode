package main

// 思路: 回溯法 组合问题 元素可以重复使用  注意比较lc78(元素不可以重复使用)
// 组合问题中 在回溯函数中需要传入起始位置，若数组中的元素可以重复使用，则向下层函数传入当前位置i
// 若数组中的元素可以重复使用，则向下层函数传入当前位置i的下一个位置，即i+1
// 这样可以避免结果集中出现重复结果

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)

	var backtrack func(nums []int, sub []int, sum int, start int)
	backtrack = func(nums []int, sub []int, sum int, start int) {
		if sum > target {
			return
		}
		if sum == target {
			item := make([]int, 0)
			item = append(item, sub...)
			res = append(res, item)
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

	backtrack(candidates, []int{}, 0, 0)
	return res
}
