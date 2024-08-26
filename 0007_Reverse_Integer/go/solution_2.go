package main

import "math"

func reverse_2(x int) int {
	res := 0
	for x != 0 {
		// 防止溢出
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}

		a := x % 10
		res = res*10 + a
		x /= 10
	}

	return res
}
