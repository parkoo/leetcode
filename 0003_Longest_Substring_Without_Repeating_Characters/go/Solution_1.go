package main

// 使用map记录并更新每个字符最后一次出现的位置
// 时间复杂度：O(n)  空间复杂度：O(min(m,n)),m表示字符集大小，n表示字符串大小

func lengthOfLongestSubstring_1(s string) int {

	start := 0
	longest_length := 0
	m := make(map[rune]int)

	for i, character := range []rune(s) {
		if index, ok := m[character]; ok && index>start {
			start = index + 1
		}

		m[character] = i

		length := i - start + 1

		if length>longest_length {
			longest_length = length
		}
	}

	return longest_length
}