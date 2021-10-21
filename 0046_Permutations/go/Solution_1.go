package main

// 回溯法
// 时间复杂度：O(n*n!)[backtract()的调用次数为n!, 每次调用从subArr复制数据到arr的时间复杂度为n]  空间复杂度：O(n)

func permute_1(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))

	var backtract func(n int, in, subArr []int)
	backtract = func(n int, in, subArr []int) {
		if len(subArr) == n {
			arr := make([]int, len(subArr))
			for i, v := range subArr {
				arr[i] = v
			}
			res = append(res, arr)
			return
		}

		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				subArr = append(subArr, in[i])
				backtract(n, in, subArr)
				subArr = subArr[:len(subArr)-1]
				used[i] = false
			}
		}
	}

	backtract(len(nums), nums, []int{})
	return res
}
