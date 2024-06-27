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
			// 如 [1, 2, 3]
			// 第一层可选 1 、2、 3
			// 当第一层选1时，第二层可选2、3
			// 当第一层选2时，第二层只可选3，
			// 当第一层选择3时，到第二层直接退出
			backtract(nums, i+1, sub)
			sub = sub[:len(sub)-1]
		}
	}

	res = append(res, []int{})
	backtract(nums, 0, []int{})

	return res
}
