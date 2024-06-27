package main

// 思路: 回溯法 组合问题 元素不可以重复使用  注意比较lc39(元素可以重复使用)
// 组合问题中，在回溯函数中需要传入起始位置，若数组中的元素可以重复使用，则向下层函数传入当前位置i
// 若数组中的元素可以重复使用，则向下层函数传入当前位置i的下一个位置，即i+1
// 这样可以避免结果集中出现重复结果

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
