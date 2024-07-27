package main

// 思路: 完全背包组合数问题 求无限物品的组合数
// 注意点：组合中不可出现重复组合  物品需在外层循环

// 注意和 lc0494的区别
// 两者实质上都是背包的组合数问题, 但lc0494属于0-1背包，每个物品只能选择一次，而本题是完全背包问题每个物品可以无限次选择
// 两者在外层循环每次确定一个物品后，都需要更新背包的状态
// 不同之处在于，0-1背包在内层循环更新背包状态时，只能依赖上一轮物品对背包状态的影响,而完全背包在内层循环更新背包状态时，可以依赖当前物品对背包状态的影响
// 对应在编码上:
// 0-1背包在内层循环时需从大到小遍历，目的是在当前物品下，只能考虑上一个物品留下的背包状态，避免当前物品被重复使用
// 而完全背包在内层循环时需从小到大遍历，目的是在当前物品下，可以使用当前物品对背包的影响，从而实现物品的重复使用

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
		for i := coin; i <= amount; i++ { // 这里需要从小到大遍历，可以理解为dp[i][j] 依赖 dp[i][j-1]
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}
