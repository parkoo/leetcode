package main

import (
	"math"
)

// 单调栈，维护一个单调递减栈
// 时间复杂度：O(n)  空间复杂度：O(n)

func trap_1(height []int) int {

	res := 0

	q := []int{}  // 以Slice实现单调递减栈

	for i:=0; i<len(height); i++ {
		for len(q)>0 && height[i]>height[q[len(q)-1]] {
			temp := height[q[len(q)-1]]
			q = q[:len(q)-1]
			if len(q)>0 {
				h := int(math.Min(float64(height[i]), float64(height[q[len(q)-1]]))) - temp
				w := i - q[len(q)-1] - 1
				res += h * w
			}
		}

		q = append(q, i)
	}

	return res
}