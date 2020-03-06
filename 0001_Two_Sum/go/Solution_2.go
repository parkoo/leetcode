package main

// 哈希表查找法，构建哈希表与哈希表查找分开
// 时间复杂度：O(n)  空间复杂度：O(n)

func twoSum_2(nums []int, target int) []int {

	// 1. 构建哈希表
	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}

	// 2. 哈希查找
	//    对于每一个元素'nums[i]',查找是否存在'target-nums[i]',注意不可重复查找同一元素
	for i:=0; i<len(nums); i++ {
		temp := target-nums[i]
		if v, ok := m[temp]; ok && v!=i {
			return []int{i, v}
		}
	}

	return nil
}