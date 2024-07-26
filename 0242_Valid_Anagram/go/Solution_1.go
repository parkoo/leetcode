package main

// 思路 ：字符串 词频统计

// 时间复杂度：O(n)  空间复杂度：O(1)

func isAnagram_1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ss, ts := []rune(s), []rune(t)
	smap, tmap := make(map[rune]int, 0), make(map[rune]int, 0)

	// 对t进行词频统计
	for _, v := range ts {
		if i, ok := tmap[v]; ok {
			tmap[v] = i + 1
		} else {
			tmap[v] = 1
		}
	}

	cnt := 0
	for i := 0; i < len(ss); i++ {
		if v, ok := smap[ss[i]]; ok {
			smap[ss[i]] = v + 1
		} else {
			smap[ss[i]] = 1
		}

		if v, ok := tmap[ss[i]]; ok && smap[ss[i]] <= v {
			cnt++
		}
	}

	return cnt == len(t)
}
