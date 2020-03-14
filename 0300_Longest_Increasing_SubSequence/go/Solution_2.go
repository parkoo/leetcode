package main

import "math"

// 贪心 + 二分查找（0ms, 击败100%）
// 贪心思想：如果要使上升子序列尽可能的长，则需要让序列上升得尽可能慢，
//         因此希望每次在上升子序列最后加上的那个数尽可能的小。
// 维护一个数组dp[], dp[length]表示递增子序列长度为length时，序列最后的那个元素值, 初始dp[1]=nums[0]
// 遍历nums，如果，nums[k]>dp[length],则dp[length+1] = nums[k], length++
//          否则，使用二分查找，在dp[1]~dp[length]中查找从左到右第一个大于nums[k]的元素dp[j],更新dp[j]=nums[k] 
// 时间复杂度：O(n*logn)  空间复杂度：O(n)

func lengthOfLIS_2(nums []int) int {
	
	if len(nums)==0 {
		return 0
	}

	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	length := 1

	for k:=1; k<len(nums); k++ {
		if nums[k]>dp[length] {
			length++
			dp[length] = nums[k]
		}else {
			// 二分查找
			l := 1
			r := length
			for l<r {
				m := l + (r-l)/2
				if dp[m]>=nums[k] {
					r = m
				}else{
					l = m+1
				}
			}

			dp[l] = nums[k]
		}
	}

	return length	
}