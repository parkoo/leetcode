package main

// 思路：贪心法  反向查找

// 时间复杂度：O(n) 空间复杂度：O(1)

func canJump_1(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	curIndex := len(nums) - 1
	for i := curIndex - 1; i >= 0; i-- {
		if nums[i] >= curIndex-i {
			if i == 0 {
				return true
			} else {
				curIndex = i
			}
		}
	}

	return false
}
