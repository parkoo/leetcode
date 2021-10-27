package main

// map妙用
// 时间复杂度：O(n)  空间复杂度：O(n)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ss, ts := []rune(s), []rune(t)
	m := make(map[rune]rune)

	for i := 0; i < len(ss); i++ {
		if v, ok := m[ss[i]]; ok {
			if ts[i] != v {
				return false
			}
		} else {
			for _, v := range m {
				if v == ts[i] {
					return false
				}
			}

			m[ss[i]] = ts[i]
		}
	}

	return true
}
