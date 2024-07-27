package main

// 思路: 完全背包组合数问题 区分组合内元素顺序
// 注意与lc0518区别，lc0518是完全背包组合数问题，但是区分组合内元素顺序

// 时间复杂度: O(target*n)    空间复杂度: O(target)

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		// 物品的遍历在内层
		// 对于每一个背包的状态，都需要从头考虑物品，达到物品的不同顺序算作不同的组合
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}
