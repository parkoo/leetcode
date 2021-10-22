package main

// 回溯法  组合问题

func subsets(nums []int) [][]int {
	res := make([][]int, 0)

	var backtrack func(start int, nums, subArr []int)
	backtrack = func(start int, nums, subArr []int) {
		for i := start; i < len(nums); i++ {
			subArr = append(subArr, nums[i])
			arr := make([]int, len(subArr))
			for i, v := range subArr {
				arr[i] = v
			}
			res = append(res, arr)
			backtrack(i+1, nums, subArr)
			subArr = subArr[:len(subArr)-1]
		}
	}

	res = append(res, []int{})
	backtrack(0, nums, []int{})
	return res
}
