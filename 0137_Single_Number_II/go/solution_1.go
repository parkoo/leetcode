package main

// 思路:  频率统计

// 时间复杂度: O(n)    空间复杂度: O(n)

func singleNumber_1(nums []int) int {
	cache := make(map[int]int)

	for _, num := range nums {
		cache[num]++
	}

	for num, cnt := range cache {
		if cnt == 1 {
			return num
		}
	}

	return -1
}
