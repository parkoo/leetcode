package main

// 双指针 滑动窗口
// 时间复杂度：O(n)  空间复杂度：O(n)

func checkInclusion(s1 string, s2 string) bool {
	ss1, ss2 := []rune(s1), []rune(s2)
	ms1, ms2 := make(map[rune]int), make(map[rune]int)

	for i := 0; i < len(ss1); i++ {
		if v, ok := ms1[ss1[i]]; ok {
			ms1[ss1[i]] = v + 1
		} else {
			ms1[ss1[i]] = 1
		}
	}

	l, r := 0, 0
	cnt := 0
	for ; r < len(ss2); r++ {
		// 扩大右边界
		if v, ok := ms2[ss2[r]]; ok {
			ms2[ss2[r]] = v + 1
		} else {
			ms2[ss2[r]] = 1
		}

		if v, ok := ms1[ss2[r]]; ok && ms2[ss2[r]] <= v {
			cnt++
		}

		// 持续缩小左边界
		for cnt == len(ss1) {
			// 检查是否满足条件
			if r-l+1 == len(ss1) {
				return true
			}

			ms2[ss2[l]] = ms2[ss2[l]] - 1
			if v, ok := ms1[ss2[l]]; ok && ms2[ss2[l]] < v {
				cnt--
			}
			l++
		}
	}

	return false
}
