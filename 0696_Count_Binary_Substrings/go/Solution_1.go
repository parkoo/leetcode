package main

// 标记法
// 时间复杂度：O(n)  空间复杂度：O(1)

func countBinarySubstrings(s string) int {
	res := 0

	ss := []rune(s)
	preLen, curLen := 0, 1
	for i := 1; i < len(ss); i++ {
		if ss[i] == ss[i-1] {
			curLen++
		} else {
			preLen = curLen
			curLen = 1
		}

		if curLen <= preLen {
			res++
		}
	}

	return res
}
