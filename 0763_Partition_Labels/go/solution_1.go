package main

// 思路: 贪心法 向外扩展边界

// 时间复杂度: O(n)    空间复杂度: O(m)  其中，m 为 s 中不重复的字符的个数

func partitionLabels(s string) []int {
	res := make([]int, 0)

	// 统计每个字符最后一次出现的位置，起始位置从1开始计算
	cache := make(map[rune]int)
	for i, char := range []rune(s) {
		cache[char] = i + 1
	}

	maxPos := 0  // 记录当前分组中可能得最远位置
	lastPos := 0 // 切分出分组时，记录上一次位置
	for i, char := range []rune(s) {
		if pos, ok := cache[char]; ok && pos > maxPos {
			maxPos = pos
		}

		if i+1 == maxPos {
			v := maxPos - lastPos
			res = append(res, v)
			lastPos = maxPos
		}
	}

	return res
}
