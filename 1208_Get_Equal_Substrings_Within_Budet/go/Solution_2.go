package main

// 双指针-滑动窗口
// 时间复杂度：O(n)  空间复杂度：O(1)

func equalSubstring_2(s string, t string, maxCost int) int {
	maxLen := 0
	curCost := 0
	curLen := 0
	// 边界思考：左开右开滑动区间 （l, r）
	for l, r := -1, 0; r < len(s); {
		if curCost+absolute(int(s[r]), int(t[r])) <= maxCost {
			curCost += absolute(int(s[r]), int(t[r]))
			curLen++
			r++
		} else {
			l++
			curCost -= absolute(int(s[l]), int(t[l]))
			curLen--
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
