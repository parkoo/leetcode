package main

// 哈希表查找法，构建哈希表的同时进行哈希查找
// 时间复杂度：O(n)  空间复杂度：O(n)

func twoSum_3(nums []int, target int) []int {

	m := make(map[int]int)

	for i, num := range nums {
		temp := target - nums[i]
		if v, ok := m[temp]; ok {
			return []int{v, i}
		}

		m[num] = i;
	}

	return nil
}