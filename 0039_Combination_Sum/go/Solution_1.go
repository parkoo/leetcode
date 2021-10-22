package main

// 回溯法 组合问题

func combinationSum_1(candidates []int, target int) [][]int {
	res := make([][]int, 0)

	var backtract func(start, target, curSum int, candidates, subArr []int)
	backtract = func(start, target, curSum int, candidates, subArr []int) {
		if curSum == target {
			arr := make([]int, len(subArr))
			for i, v := range subArr {
				arr[i] = v
			}
			res = append(res, arr)

			return
		}

		for i := start; i < len(candidates); i++ {
			if curSum+candidates[i] <= target {
				subArr = append(subArr, candidates[i])
				backtract(i, target, curSum+candidates[i], candidates, subArr)
				subArr = subArr[:len(subArr)-1]
			}
		}
	}

	backtract(0, target, 0, candidates, []int{})
	return res
}
