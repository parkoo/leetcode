package main

// 回溯法  组合问题

func subsets(nums []int) [][]int {
	res := make([][]int, 0)

	var backtract func(nums []int, start int, sub []int)
	backtract = func(nums []int, start int, sub []int) {
		for i := start; i < len(nums); i++ {
			sub = append(sub, nums[i])
			item := make([]int, 0)
			item = append(item, sub...)
			res = append(res, item)

			// 当前层选用第i个数字时，下一层的起始位置从i+1开始
			backtract(nums, i+1, sub)
			sub = sub[:len(sub)-1]
		}
	}

	res = append(res, []int{})
	backtract(nums, 0, []int{})

	return res
}
