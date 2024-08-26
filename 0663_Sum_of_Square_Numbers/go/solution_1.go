package main

import "math"

// 思路: 双指针夹逼

// 时间复杂度: O(sqrt(c))    空间复杂度: O(1)

func judgeSquareSum(c int) bool {
	l, r := 0, int(math.Sqrt(float64(c)))

	for l <= r {
		sum := l*l + r*r
		if sum == c {
			return true
		}

		if sum > c {
			r--
		} else {
			l++
		}
	}

	return false
}
