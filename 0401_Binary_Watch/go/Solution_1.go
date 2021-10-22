package main

import (
	"fmt"
)

// 回溯法 N个位置中选K个位置的组合
// 时间复杂度：O(n!)  空间复杂度：O(n)

func readBinaryWatch_1(turnedOn int) []string {
	res := make([]string, 0)

	var backtract func(start, cnt, k, n int, subArr []int)
	backtract = func(start, cnt, k, n int, subArr []int) {
		if cnt == k {
			hour := subArr[9]*8 + subArr[8]*4 + subArr[7]*2 + subArr[6]
			min := subArr[5]*32 + subArr[4]*16 + subArr[3]*8 + subArr[2]*4 + subArr[1]*2 + subArr[0]

			if hour < 12 && min < 60 {
				time := fmt.Sprintf("%d:%02d", hour, min)
				res = append(res, time)
			}

			return
		}

		for i := start; i <= n; i++ {
			subArr[i] = 1
			backtract(i+1, cnt+1, k, n, subArr)
			subArr[i] = 0
		}
	}

	subArr := make([]int, 10)
	backtract(0, 0, turnedOn, 9, subArr)
	return res
}
