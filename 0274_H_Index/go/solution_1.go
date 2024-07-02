package main

// 思路:  计数统计

// 时间复杂度: O(n)    空间复杂度: O(n)

func hIndex_1(citations []int) int {
	n := len(citations)

	cache := make(map[int]int) // 记录 [引用次数]->[该引用次数的文章数]
	for _, citation := range citations {
		if citation >= n {
			citation = n // 因为H值不能超过文章的总篇数， 大于总片数的引用值按照总片数算
		}
		cache[citation]++
	}

	total := 0 // 记录大于等于某个引用次数的文章的总篇数
	// 从大到小遍历，方便记录total
	for citation := n; citation >= 0; citation-- {
		total += cache[citation] // total为大于等于citation的文章的总篇数

		if total >= citation {
			return citation
		}
	}

	return 0
}
