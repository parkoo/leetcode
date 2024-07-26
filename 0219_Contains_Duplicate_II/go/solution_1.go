package main

// 思路: map 记录位置

// 时间复杂度: O(n)    空间复杂度: O(n)

func containsNearbyDuplicate(nums []int, k int) bool {

	indexCahce := make(map[int]int)
	for i, num := range nums {
		lastIndex, ok := indexCahce[num]
		if ok && i-lastIndex <= k {
			return true
		}

		indexCahce[num] = i
	}

	return false
}
