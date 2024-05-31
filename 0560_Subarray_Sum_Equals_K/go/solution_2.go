package main

// 思路: 前缀和 + 哈希表
// 对于每个前缀和preSum[i]，都检查是否存在一个早先的前缀和preSum[j]，使得它们的差等于k (preSum[i]-preSum[j]==k)
// 如果存在，就找到了一个和为k的子数组 [j,...,i]

// 时间复杂度: O(n)    空间复杂度: O(n)

func subarraySum_2(nums []int, k int) int {
	res := 0

	cache := make(map[int]int) // 记录每个前缀和出现的次数
	cache[0] = 1               // 初始化cache 处理前缀和preSum[i]恰好等于k的情况

	preSum := 0 // 前缀和
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]

		if cnt, ok := cache[preSum-k]; ok {
			res += cnt
		}

		cache[preSum]++
	}

	return res
}
