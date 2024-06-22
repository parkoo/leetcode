package main

// 思路: 贪心法  向外扩展边界

// 时间复杂度：O(n) 空间复杂度：O(1)

func canJump_2(nums []int) bool {
	maxDistance := 0 // 可以到达的最远距离，实区间
	for i, num := range nums {
		if i > maxDistance { // 此时已经到达最远距离之外
			return false
		}
		if i+num > maxDistance { // 在可达范围内尝试扩展距离
			maxDistance = i + num
		}
	}

	return maxDistance >= len(nums)-1 // 最远距离是否超过数组边界
}
