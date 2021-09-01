package main

import "math/rand"

// 快排partition思想 若某次partition得到的位置p恰好是k则直接返回 否则根据p与k的大小,选择向一边进行partition
// 时间复杂度：O(n)  空间复杂度；O(lgn)
func findKthLargest_1(nums []int, k int) int {
	return QuickSort(nums, 0, len(nums)-1, k-1)
}

func QuickSort(nums []int, l, r, k int) int {
	if l >= r {
		return nums[l]
	}

	p := partition(nums, l, r)
	if p == k {
		return nums[p]
	} else if p > k {
		return QuickSort(nums, l, p-1, k)
	} else {
		return QuickSort(nums, p+1, r, k-p)
	}
}

func partition(nums []int, l, r int) int {
	p := rand.Intn(r-l+1) + l
	nums[l], nums[p] = nums[p], nums[l]

	i, j := l+1, r
	for {
		for i <= r && nums[i] > nums[l] {
			i++
		}
		for j >= l+1 && nums[j] < nums[l] {
			j--
		}

		if i > j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	nums[l], nums[j] = nums[j], nums[l]
	return j
}
