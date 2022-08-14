package main

// 前缀和
// 时间复杂度：O(n)  空间复杂度：O(n)

func maxScore_2(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	}

	presum := make([]int, len(s)+1)
	for i := 1; i <= len(s); i++ {
		presum[i] = presum[i-1] + int(s[i-1]-'0')
	}

	for i := 0; i < len(s)-1; i++ {
		temp := (i + 1 - presum[i+1]) + (presum[len(s)] - presum[i+1])
		if temp > res {
			res = temp
		}
	}

	return res
}
