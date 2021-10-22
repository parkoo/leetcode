package main

import (
	"sort"
)

// 回溯法
// 先排序，然后在添加一个元素时，判断这个元素是否等于前一个元素，如果等于，并且前一个元素还未访问，那么就跳过这个元素
// 时间复杂度：O(n*n!)[backtract()的调用次数为n!, 每次调用从subArr复制数据到arr的时间复杂度为n]  空间复杂度：O(n)

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))

	var backtrack func(n int, in, subArr []int)
	backtrack = func(n int, in, subArr []int) {
		if len(subArr) == n {
			arr := make([]int, n)
			for i, v := range subArr {
				arr[i] = v
			}
			res = append(res, arr)
			return
		}

		for i := 0; i < n; i++ {
			if !used[i] {
				// 要保证在同一层中不使用相同的数字
				// 若某一个元素和上一个元素相等,并且上一个元素没有被使用过，说明它俩在同一层;
				// 若某一个元素和上一个元素相等,并且上一个元素被使用过，说明这个元素在上一个元素的子树中
				if i > 0 && in[i] == in[i-1] && used[i-1] {
					continue
				}
				used[i] = true
				subArr = append(subArr, in[i])
				backtrack(n, in, subArr)
				subArr = subArr[:len(subArr)-1]
				used[i] = false
			}
		}
	}

	sort.Slice(nums, func(i int, j int) bool { return nums[i] < nums[j] })
	backtrack(len(nums), nums, []int{})
	return res
}
