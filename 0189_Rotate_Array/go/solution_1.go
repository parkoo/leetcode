package main

// 思路: 整体翻转 + 局部翻转 实现数组旋转效果

// 时间复杂度: O(n)    空间复杂度: O(1)

func rotate_1(nums []int, k int) {
	k = k % len(nums)

	// 整体翻转
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	// 翻转前k个数
	for i, j := 0, k-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	// 翻转剩下的数
	for i, j := k, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
