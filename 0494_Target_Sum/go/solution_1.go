package main

// 思路: 0-1背包组合数问题
// 准备两个背包，一个背包package_a存放标记为正的元素，另一个背包package_b存放标记为负的元素, 则可知 package_a - package_b = target
// 设nums的元素和为sum, 可以列出方程
// package_a - package_b = target
// package_a + package_b = sum
// 则 package_a = (target + sum)/2。 所以根据题意给的target和sum，可以求出package_a的值。
// 那这道题就可以转化为：给定一个大小为package_a的背包，有多少种组合方式能把背包装满, 即0-1背包组合数问题

// 注意和 lc0518的区别
// 两者实质上都是背包的组合数问题, 但lc0518属于完全背包，每个物品每个物品可以无限次选择，而本题是只能选择一次
// 两者在外层循环每次确定一个物品后，都需要更新背包的状态
// 不同之处在于，0-1背包在内层循环更新背包状态时，只能依赖上一轮物品对背包状态的影响,而完全背包在内层循环更新背包状态时，可以依赖当前物品对背包状态的影响
// 对应在编码上:
// 0-1背包在内层循环时需从大到小遍历，目的是在当前物品下，只能考虑上一个物品留下的背包状态，避免当前物品被重复使用
// 而完全背包在内层循环时需从小到大遍历，目的是在当前物品下，可以使用当前物品对背包的影响，从而实现物品的重复使用

// 时间复杂度: O(n×(sum−target))    空间复杂度: O(sum−target)

func findTargetSumWays(nums []int, target int) int {
	// 计算背包并确保背包有效
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (target+sum)%2 == 1 {
		return 0
	}
	cap := (target + sum) / 2
	if cap < 0 {
		return 0
	}

	// 0-1背包组合数问题
	dp := make([]int, cap+1)
	dp[0] = 1

	for _, num := range nums {
		// 0-1背包 从大到小遍历
		for i := cap; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}

	return dp[cap]
}
