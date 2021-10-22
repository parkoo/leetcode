package main

// 动态规划  组合问题

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)

	var backtrack func(start, k, n, curSum int, subArr []int)
	backtrack = func(start, k, n, curSum int, subArr []int) {
		if len(subArr) == k {
			if curSum == n {
				arr := make([]int, k)
				for i, v := range subArr {
					arr[i] = v
				}
				res = append(res, arr)
			}

			return
		}

		// 剪枝1：i<=9  -->  i <= 9-(k-len(subArr))+1
		for i := start; i <= 9-(k-len(subArr))+1; i++ {
			if curSum+i <= n { // 剪枝2
				subArr = append(subArr, i)
				backtrack(i+1, k, n, curSum+i, subArr)
				subArr = subArr[:len(subArr)-1]
			}
		}
	}

	backtrack(1, k, n, 0, []int{})
	return res
}
