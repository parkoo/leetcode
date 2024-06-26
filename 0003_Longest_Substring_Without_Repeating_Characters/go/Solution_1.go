package main

// 思路：使用map记录并更新每个字符最后一次出现的位置

// 时间复杂度：O(n)  空间复杂度：O(min(m,n)),m表示字符集大小，n表示字符串大小

func lengthOfLongestSubstring_1(s string) int {
	res := 0

	cache := make(map[rune]int)
	left := -1 // (left, right] 左开右闭区间为不重复内容
	for right := 0; right < len(s); right++ {
		cur := []rune(s)[right]
		lastPos, ok := cache[cur]
		if ok && lastPos > left {
			left = lastPos
		}

		// 每次都尝试更新最大长度
		curLen := right - left // (left, right] 的长度
		if curLen > res {
			res = curLen
		}

		cache[cur] = right
	}

	return res
}
