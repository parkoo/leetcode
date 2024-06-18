package main

// 双指针 滑动窗口, 滑动窗口类型的问题中都会有两个指针，一个用于"延伸"现有窗口的 r 指针，和一个用于"收缩"窗口的 l 指针。
// 在任意时刻，只有一个指针运动，而另一个保持静止

// 时间复杂度：O(s + t)  空间复杂度：O(c)  其中, c为字符集的大小

func minWindow(s string, t string) string {
	res := ""
	ss, ts := []rune(s), []rune(t)
	smap, tmap := make(map[rune]int, 0), make(map[rune]int, 0)

	// 对t进行词频统计
	for _, v := range ts {
		tmap[v]++
	}

	l := len(ss) + 1 // 初始化s中包含t的最小长度
	cnt := 0         // 记录窗口[i,j]中记录的t中字符的个数
	for i, j := 0, 0; j < len(ss); j++ {
		// 扩大窗口右边界
		smap[ss[j]]++
		if v, ok := tmap[ss[j]]; ok && smap[ss[j]] <= v {
			cnt++
		}

		// 缩小窗口左边界
		for cnt == len(ts) {
			if j-i+1 < l {
				l = j - i + 1
				res = s[i : j+1]
			}

			smap[ss[i]]--
			if v, ok := tmap[ss[i]]; ok && smap[ss[i]] < v {
				cnt--
			}
			i++
		}
	}

	return res
}
