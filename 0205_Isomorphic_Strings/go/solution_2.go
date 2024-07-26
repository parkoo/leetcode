package main

// 思路： map妙用 同lc090

// 时间复杂度：O(n)  空间复杂度：O(n)

func isIsomorphic_2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ss, tt := []rune(s), []rune(t)
	stMap := make(map[rune]rune)      // 记录s中的字符与t中对应位置处的字符的映射
	tCache := make(map[rune]struct{}) // 记录t中出现过的字符

	for i := 0; i < len(ss); i++ {
		sVal, tVal := ss[i], tt[i]
		val, ok := stMap[sVal]
		if ok {
			if val != tVal {
				return false
			}

		} else {
			if _, ok := tCache[tVal]; ok {
				return false
			}

			stMap[sVal] = tVal
			tCache[tVal] = struct{}{}
		}
	}

	return true
}
