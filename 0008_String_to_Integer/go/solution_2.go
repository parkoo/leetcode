package main

import (
	"math"
)

// 时间复杂度：O(m),为有效数字的长度  空间复杂度：O(n)，数组strArr所占

func myAtoi_2(str string) int {
	n := len(str)

	// 处理前置空格
	i := 0
	for i < n && str[i] == ' ' {
		i++
	}

	// 判断符号
	flag := 1 // flag=1表示正数，flag=-1表示负数
	if i < n && (str[i] == '-' || str[i] == '+') {
		if str[i] == '-' {
			flag = -1
		}
		i++
	}

	res := 0
	for i < n {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		cur := int(str[i] - '0')

		// 处理溢出
		if res > (math.MaxInt32-cur)/10 {
			if flag == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		res = res*10 + cur
		i++
	}

	return res * flag

}
