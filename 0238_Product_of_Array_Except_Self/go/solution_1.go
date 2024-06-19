package main

// 思路: 左右分别统计
// 分别统计当前元素左边元素的乘积和右边元素的乘积, 左右乘积相乘即是除自身之外的乘积

// 时间复杂度: O(n)    空间复杂度: O(n)

func productExceptSelf_1(nums []int) []int {
	res := make([]int, len(nums))

	// 分别统计当前元素左边的的乘积和右边的乘积
	leftProduct, rightProduct := make([]int, len(nums)), make([]int, len(nums))
	leftProduct[0], rightProduct[len(nums)-1] = 1, 1

	for i := 1; i < len(nums); i++ {
		leftProduct[i] = nums[i-1] * leftProduct[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		rightProduct[i] = nums[i+1] * rightProduct[i+1]
	}

	// 左右乘积相乘即是除自身之外的乘积
	for i := 0; i < len(nums); i++ {
		res[i] = leftProduct[i] * rightProduct[i]
	}

	return res
}
