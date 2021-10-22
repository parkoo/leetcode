package main

import (
	"sort"
)

// 回溯法  组合问题

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))

	var backtrack func(start int, nums, subArr []int)
	backtrack = func(start int, nums, subArr []int) {
		for i := start; i < len(nums); i++ {
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}

			subArr = append(subArr, nums[i])
			used[i] = true
			arr := make([]int, len(subArr))
			for i, v := range subArr {
				arr[i] = v
			}
			res = append(res, arr)
			backtrack(i+1, nums, subArr)
			subArr = subArr[:len(subArr)-1]
			used[i] = false
		}
	}

	sort.Slice(nums, func(i int, j int) bool { return nums[i] < nums[j] })
	res = append(res, []int{})
	backtrack(0, nums, []int{})
	return res
}
