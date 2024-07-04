package main

// 思路：双指针  滑动窗口

// 时间复杂度：O(n)  空间复杂度：O(1)

func minSubArrayLen_1(target int, nums []int) int {
	cnt := len(nums) + 1
	i, j := 0, -1 // 滑动串口为[i,j], j<i时表示窗口中无元素，j=i时表示窗口中有一个元素
	sum := 0
	for i < len(nums) {
		if j < len(nums)-1 && sum < target { // 防止j越界
			j++
			sum += nums[j]
		} else {
			sum -= nums[i]
			i++
		}

		if sum >= target && j-i+1 < cnt {
			cnt = j - i + 1
		}
	}

	if cnt == len(nums)+1 {
		cnt = 0
	}
	return cnt
}
