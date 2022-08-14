package main

// 前缀和 + 单调栈 + 贪心
// 类似 0962_Maximum_Width_Ramp

func longestWPI(hours []int) int {
	res := 0
	presum := make([]int, len(hours)+1)
	for i := 1; i < len(presum); i++ {
		presum[i] = presum[i-1] + transHour(hours[i-1])
	}

	stack := make([]int, 0)
	for i, sum := range presum {
		if len(stack) == 0 || sum < presum[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}

	for j := len(presum) - 1; j >= 0; j-- {
		for len(stack) > 0 && presum[j] > presum[stack[len(stack)-1]] {
			cur := j - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if cur > res {
				res = cur
			}
		}
	}

	return res
}

func transHour(hour int) int {
	if hour > 8 {
		return 1
	}
	return -1
}
