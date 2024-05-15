package main

// 回溯法 组合问题

func combinationSum_1(candidates []int, target int) [][]int {
	res := make([][]int, 0)

	var backtract func(start, curSum int, candidates, subArr []int)
	backtract = func(start, curSum int, candidates, subArr []int) {
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
				// start传入i, 保证下层节点与当前节点从同一个位置开始，避免结果重复 (如：[2,2,3]、[2,3,2]、[3,2,2])
				backtract(i, curSum+candidates[i], candidates, subArr)
				subArr = subArr[:len(subArr)-1]
			}
		}
	}

	backtract(0, 0, candidates, []int{})
	return res
}
