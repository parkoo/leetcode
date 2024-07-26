package main

// 思路 ：字符串 词频统计 同lc0383

// 时间复杂度：O(n)  空间复杂度：O(1)

func isAnagram_2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ss, tt := []rune(s), []rune(t)
	sFreq := make(map[rune]int)
	for _, c := range ss {
		sFreq[c]++
	}

	for _, c := range tt {
		if cnt, ok := sFreq[c]; !ok || cnt == 0 {
			return false
		}

		sFreq[c]--
	}

	return true
}
