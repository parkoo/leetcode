package main

// 使用map记录并更新每个字符最后一次出现的位置
// 时间复杂度：O(n)  空间复杂度：O(min(m,n)),m表示字符集大小，n表示字符串大小

func lengthOfLongestSubstring_1(s string) int {
	start := 0
	longestLength := 0
	locMap := make(map[rune]int)
	for i, character := range []rune(s) {
		if index, ok := locMap[character]; ok && index+1 > start {
			start = index + 1
		}
		length := i - start + 1
		if length > longestLength {
			longestLength = length
		}
		locMap[character] = i
	}

	return longestLength
}
