package main

// 思路: 动态规划

// 时间复杂度: O(n)    空间复杂度: O(n)

func numberOfArithmeticSlices_1(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	// 定义dp[i] = k, 表示 ‘由nums[i]结尾’ 的等差数列的个数为k

	// 若 nums[i] - nums[i-1] == nums[i-1] - nums[i-2]
	// 则属于dp[i]等差数列（表示由nums[i]结尾的等差数列）由两部分构成：
	// 1. 由属于dp[i-1]的每个等差数列后边再加一个nums[i]可以得到一个个属于dp[i]等差数列
	// 2. 新增的一个[nums[i-2],nums[i-1],nums[i]] 等差数列 （因为最少三个数构成一个等差数列）
	// 故 dp[i] = dp[i-1] + 1

	//  若 nums[i] - nums[i-1] != nums[i-1] - nums[i-2]
	//  则以nums[i]结尾的等差数列个数为0， 即 dp[i] = 0

	// 如：
	// [1,2,3]     -> dp[2] = 1 (表示以nums[2],即数字3结尾的等差数列的个数) -> [1,2,3])
	// [1,2,3,4]   -> dp[3] = 2 (表示以nums[3],即数字4结尾的等差数列的个数) -> [1,2,3,4] + [2,3,4]
	// [1,2,3,4,5] -> dp[4] = 3 (表示以nums[4],即数字5结尾的等差数列的个数) -> [1,2,3,4,5] [2,3,4,5] + [3,4,5]

	// 而所有的等差数列的个数为 res := dp[2] + dp[3] + ... + dp[n-1]

	res := 0
	dp := make([]int, n)
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}

		res += dp[i]
	}

	return res
}
