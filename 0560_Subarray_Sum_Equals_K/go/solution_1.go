package main

// 思路: 暴力枚举

// 时间复杂度: O(n^2)    空间复杂度: O(1)

func subarraySum_1(nums []int, k int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
	}

	return res
}
