package main

// 思路： 贪心 + 二分查找（0ms, 击败100%）

// 贪心思想：如果要使上升子序列尽可能的长，则需要让序列上升得尽可能慢，
//         因此希望每次在上升子序列最后加上的那个数尽可能的小。

// 维护一个数组d[], d[length] = k 表示长度为length的递增子序列的最后一个元素最小为k, 初始dp[1]=nums[0]
// 若 l = 3, 对于 [2, 3, 6, 5], 则需要维护d[3] = 5
// 可知 d 是一个递增数组, 即 i > j, 则 d[i] > d[j]
// 遍历nums，如果，nums[i]>d[length],则d[length+1] = nums[i], length++
//          否则，使用二分查找，在dp[1]~dp[length]中查找从左到右第一个大于nums[i]的元素d[j],更新dp[j]=nums[i]

// 时间复杂度：O(n*logn)  空间复杂度：O(n)

func lengthOfLIS_2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 定义 d[l] = k 表示长度为l的递增子序列的最后一个元素最小为k
	// 若 l = 3, 对于 [2, 3, 6, 5], 则需要维护d[3] = 5
	// 可知 d 是一个递增数组, 即 i > j, 则 d[i] > d[j]
	d := make([]int, n+1)

	maxLen := 1 // 维护当前最长递增子序列的长度
	d[maxLen] = nums[0]

	for i := 1; i < n; i++ {
		if nums[i] > d[maxLen] {
			maxLen++
			d[maxLen] = nums[i]

		} else {
			index := binarySearch(d, maxLen, nums[i])
			d[index] = nums[i]
		}
	}

	return maxLen
}

func binarySearch(arr []int, maxLen int, target int) int {
	left, right := 1, maxLen
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] >= target {
			right = mid

		} else {
			left = mid + 1
		}
	}

	return left
}
