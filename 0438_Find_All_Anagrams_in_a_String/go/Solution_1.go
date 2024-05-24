package main

// 双指针 滑动窗口
// 时间复杂度：O(n)  空间复杂度：O(1)

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	ss, ps := []rune(s), []rune(p)
	smap, pmap := make(map[rune]int, 0), make(map[rune]int, 0)

	for _, v := range ps {
		pmap[v]++
	}

	i, j := 0, 0
	cnt := 0 // 记录滑动窗口中保存的p中字符的个数
	for ; j < len(ss); j++ {
		smap[ss[j]]++

		if v, ok := pmap[ss[j]]; ok && smap[ss[j]] <= v {
			cnt++
		}

		for cnt == len(ps) {
			if j-i+1 == len(ps) {
				res = append(res, i)
			}

			// 缩小左边界
			if _, ok := pmap[ss[i]]; ok && smap[ss[i]] <= pmap[ss[i]] {
				cnt--
			}
			smap[ss[i]] = smap[ss[i]] - 1
			i++
		}
	}

	return res
}
