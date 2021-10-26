package main

// 字符串 词频统计
// 时间复杂度：O(n)  空间复杂度：O(1)

func longestPalindrome(s string) int {
	// 词频统计
	smap := make(map[rune]int, 0)
	for _, v := range []rune(s) {
		if i, ok := smap[v]; ok {
			smap[v] = i + 1
		} else {
			smap[v] = 1
		}
	}

	cnt := 0
	for _, v := range smap {
		if v%2 == 0 {
			cnt += v
		} else {
			cnt += v - 1
		}
	}

	// 中间加一个字符
	if cnt < len(s) {
		cnt++
	}

	return cnt
}
