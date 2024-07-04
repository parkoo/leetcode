package main

// 思路：双指针  滑动窗口  同solution_1

// 时间复杂度：O(n)  空间复杂度：O(1)

func minSubArrayLen_2(target int, nums []int) int {
	n := len(nums)
	res := n + 1

	curSum := 0
	left := -1
	for right := 0; right < n; right++ {
		// 扩展右边界
		curSum += nums[right]

		// 持续缩小左边界直到窗口不满足条件
		for curSum >= target {
			curLen := right - left
			if curLen < res {
				res = curLen
			}

			left++
			leftVal := nums[left]
			curSum -= leftVal
		}
	}

	if res == n+1 {
		res = 0
	}
	return res
}
