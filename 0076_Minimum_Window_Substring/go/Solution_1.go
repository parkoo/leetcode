package main

// 双指针 滑动窗口
// 时间复杂度：O(n)  空间复杂度：O(1)

func minWindow(s string, t string) string {
	res := ""
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

	l := len(ss) + 1 // 初始化s中包含t的最小长度
	i, j := 0, 0
	cnt := 0 // 记录窗口[i,j]中记录的t中字符的个数
	for ; j < len(ss); j++ {
		// 扩大窗口右边界
		if v, ok := smap[ss[j]]; ok {
			smap[ss[j]] = v + 1
		} else {
			smap[ss[j]] = 1
		}

		if v, ok := tmap[ss[j]]; ok && smap[ss[j]] <= v {
			cnt++
		}

		// 缩小窗口左边界
		for cnt == len(ts) {
			if j-i+1 < l {
				l = j - i + 1
				res = string(ss[i : j+1])
			}

			smap[ss[i]] = smap[ss[i]] - 1
			if v, ok := tmap[ss[i]]; ok && smap[ss[i]] < v {
				cnt--
			}
			i++
		}
	}

	return res
}
