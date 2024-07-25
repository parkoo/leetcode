package main

// 思路：组合问题　回溯

// 时间复杂度：O(k*Cₙᵏ)  空间复杂度：O(Cₙᵏ)  其中,k表示需要将每个长度为k的组合都添加到输出中

func combine_2(n int, k int) [][]int {
	res := make([][]int, 0)

	var backtract func(cnt, start int, sub []int)
	backtract = func(cnt, start int, sub []int) {
		if cnt == k {
			item := make([]int, 0)
			item = append(item, sub...)
			res = append(res, item)
			return
		}

		for i := start; i <= n; i++ {
			sub = append(sub, i)
			backtract(cnt+1, i+1, sub)
			sub = sub[:len(sub)-1]
		}

	}

	backtract(0, 1, []int{})

	return res
}
