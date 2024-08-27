package main

// 思路：双指针-滑动窗口

// 时间复杂度：O(n)  空间复杂度：O(1)

func equalSubstring_3(s string, t string, maxCost int) int {
	ss, ts := []rune(s), []rune(t)

	curCost := 0
	maxLen := 0
	l := -1
	for r := 0; r < len(s); r++ {
		sc, tc := ss[r], ts[r]
		cost := getCost(sc, tc)
		curCost += cost

		for curCost > maxCost {
			l++
			curCost -= getCost(ss[l], ts[l])
		}

		curLen := r - l
		if curLen > maxLen {
			maxLen = r - l
		}
	}

	return maxLen
}

func getCost(a, b rune) int {
	res := int(a) - int(b)
	if res < 0 {
		res = -res
	}
	return res
}
