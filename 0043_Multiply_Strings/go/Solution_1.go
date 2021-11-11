package main

import (
	"strconv"
)

// 时间复杂度：O(mn)  空间复杂度：O(m+n)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	resArr := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		x := int(num1[i]) - '0'
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j]) - '0'
			resArr[i+j+1] += x * y
		}
	}

	for i := m + n - 1; i > 0; i-- {
		resArr[i-1] += resArr[i] / 10
		resArr[i] %= 10
	}

	res := ""
	i := 0
	if resArr[0] == 0 {
		i = 1
	}
	for ; i < len(resArr); i++ {
		res = res + strconv.Itoa(resArr[i])
	}

	return res
}
