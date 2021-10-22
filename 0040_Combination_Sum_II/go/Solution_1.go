package main

import (
	"sort"
)

// 回溯法 组合问题

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(candidates))

	var backtract func(start, curSum, target int, candidates, subArr []int)
	backtract = func(start, curSum, target int, candidates, subArr []int) {
		if curSum == target {
			arr := make([]int, len(subArr))
			for i, v := range subArr {
				arr[i] = v
			}
			res = append(res, arr)
			return
		}

		for i := start; i < len(candidates); i++ {
			// 去重
			if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
				continue
			}

			if curSum+candidates[i] <= target {
				used[i] = true
				subArr = append(subArr, candidates[i])
				backtract(i+1, curSum+candidates[i], target, candidates, subArr)
				subArr = subArr[:len(subArr)-1]
				used[i] = false
			}
		}
	}

	sort.Slice(candidates, func(i int, j int) bool { return candidates[i] < candidates[j] })
	backtract(0, 0, target, candidates, []int{})
	return res
}
