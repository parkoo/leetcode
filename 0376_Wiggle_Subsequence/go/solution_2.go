package main

// 思路:  子序列问题  全局性考虑

// 时间复杂度: O(n)    空间复杂度: O(1)

func wiggleMaxLength_2(nums []int) int {
	res := 0
	n := len(nums)
	if n == 0 {
		return res
	}

	// up 表示考虑到当前位置i, 考虑将nums[i]放入摆动序列的上升点（可以选择放入或者不放入），全局的最长摆动序列的长度
	// down 表示考虑到当前位置i, 考虑将nums[i]放入摆动序列的下降点（可以选择放入或者不放入），全局的最长摆动序列的长度
	up, down := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		}
		if nums[i] < nums[i-1] {
			down = up + 1
		}
	}

	res = up
	if down > up {
		res = down
	}

	return res
}
