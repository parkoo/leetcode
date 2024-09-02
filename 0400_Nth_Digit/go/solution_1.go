package main

import "math"

// 思路: 模拟
// 对于长度为 len 的数字的范围为 [10^(len−1), 10^len -1]（共 9 *10^(len−1)个），总长度为：L=len*9*10^(len−1)

// https://leetcode.cn/problems/nth-digit/solutions/1129553/gong-shui-san-xie-jian-dan-mo-ni-ti-by-a-w5wl/

// 时间复杂度: O(lgn)    空间复杂度: O(1)

func findNthDigit_1(n int) int {
	len := 1
	for len*9*int(math.Pow(float64(10), float64(len-1))) < n {
		n -= len * 9 * int(math.Pow(float64(10), float64(len-1)))
		len++
	}

	s := int64(math.Pow(float64(10), float64(len-1)))
	s += int64(n/len - 1)
	n -= len * (n / len)
	if n == 0 {
		return int(s % 10)

	} else {
		return int(int(s+1) / int(math.Pow(float64(10), float64(len-n))) % 10)
	}
}
