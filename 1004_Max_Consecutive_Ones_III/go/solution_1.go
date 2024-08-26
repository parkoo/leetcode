package main

// 思路: 滑动窗口
// 对于数组 nums 的区间 [left,right] 而言，只要它包含不超过 k 个 0，我们就可以根据它构造出一段满足要求，并且长度为 right−left+1 的区间

// 时间复杂度: O(n)    空间复杂度: O(1)

func longestOnes_1(nums []int, k int) int {
	maxLen := 0
	cnt := 0
	l := -1
	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			cnt++
		}

		// 注意这里必须是cnt > k, 保证0个个数超额后再缩小左边界矫正，而不是 cnt==k
		// [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k=3
		for cnt > k {
			l++
			if nums[l] == 0 {
				cnt--
			}
		}

		// 经过上述右边界的扩大和左边界的缩小后，cnt<=k
		curLen := r - l
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}
