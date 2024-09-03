package main

// 思路: 原地置换

// 时间复杂度: O(n)    空间复杂度: O(1)

func findRepeatDocument_1(documents []int) int {
	n := len(documents)
	for i, doc := range documents {
		for doc >= 0 && doc < n && documents[doc] != doc {
			documents[i], documents[doc] = documents[doc], documents[i]
		}
	}

	for i, doc := range documents {
		if i != doc {
			return doc
		}
	}

	return -1
}
