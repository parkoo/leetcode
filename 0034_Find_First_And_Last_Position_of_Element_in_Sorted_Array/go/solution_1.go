package main

// 思路: 二分查找, 定位左边界后再定位右边界

// 时间复杂度: O(log(n))    空间复杂度: O(1)

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}

	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// 不存在
	if l >= len(nums) || nums[l] != target {
		return res
	}

	// 向右寻找右边界
	end := l
	for end < len(nums) && nums[end] == target {
		end++
	}
	res[0], res[1] = l, end-1

	return res
}
