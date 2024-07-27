package main

// 思路: 背包问题 求无限物品的组合数
// 注意点：组合中不可出现重复组合  物品需在外层循环

// 时间复杂度: O(n*amount)    空间复杂度: O(amount)

func change(amount int, coins []int) int {

	// coin在外层，可以确保不会计算出重复的组合
	// 取出一个硬币，向所有能放入的背包中放入，所以硬币是可重复使用的
	// 而硬币的取出又是有顺序的，所以不会出现重复组合
	// 如：对于coins=[1,2], amount=3
	// 出现的组合为：
	// 3 = 1 + 1 + 1
	// 3 = 1 + 2
	// 不会出现：
	// 3 = 2 + 1

	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}
