package main

// 思路: 枚举

// 如果一个长度为n的字符串s,可以由它的一个长度为preLen的子串pre重复多次构成，那么：
// n 一定是 preLen的倍数, pre 一定是 s 的前缀；
// 对于任意的位置i, 有 s[i]=s[i−preLen]

// 时间复杂度: O(n^2)    空间复杂度: O(1)

func repeatedSubstringPattern(s string) bool {
	n := len(s)

	// preLen 表示前缀的长度，从1开始增大, 暴力对preLen进行尝试
	for preLen := 1; preLen*2 <= n; preLen++ {
		// n一定是前缀长度的整数倍
		if n%preLen != 0 {
			continue
		}

		match := true
		for j := preLen; j < n; j++ {
			if s[j] != s[j-preLen] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}

	return false
}
