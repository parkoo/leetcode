package main

// 暴力法
// 时间复杂度：O(nm)  空间复杂度：O(1)  （n为字符串的长度，m为最长子串的长度）

func equalSubstring_1(s string, t string, maxCost int) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		curLen := 0
		cost := 0
		for j := i; j < len(s); j++ {
			cost += absolute(int(s[j]), int(t[j]))
			if cost > maxCost {
				break
			}
			curLen++
		}
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}

func absolute(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}
