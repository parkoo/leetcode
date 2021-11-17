package main

// map统计存在性及去重
// 时间复杂度：O(n)  空间复杂度：O(n)

func longestConsecutive(nums []int) int {
	// map用于去重并记录数字的存在性
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}

	res := 0
	for num := range m {
		// 若前一个数字存在则 从前一个数字计算长度会更长，直接contine, 避免重复计算长度
		if m[num-1] {
			continue
		}

		curNum := num
		cnt := 0
		for m[curNum] {
			cnt++
			curNum++
		}
		if cnt > res {
			res = cnt
		}
	}

	return res
}
